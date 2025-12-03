# MinIO Best Practices Guide with minio-go SDK

This document outlines best practices for integrating MinIO (S3-compatible object storage) into Go applications using the official `github.com/minio/minio-go/v7` SDK. It draws from the official MinIO Go Client SDK documentation, API reference, quickstart examples, and community recommendations to ensure secure, performant, and scalable object storage operations. These guidelines are tailored for use with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus, as well as dependencies like `pgx`, `mongo-go-driver`, `amqp091-go`, `go-redis`, `sony/gobreaker`, `avast/retry-go`, and `ulule/limiter`. The SDK supports Go 1.21+ and is compatible with any S3-compatible storage, including MinIO servers (e.g., play.min.io for testing).

## Installation and Setup

- Install the MinIO Go Client SDK using Go modules:
  ```bash:disable-run
  go get github.com/minio/minio-go/v7
  go get github.com/minio/minio-go/v7/pkg/credentials  # For advanced credential management
  ```
- Import the necessary packages:
  ```go
  import (
      "context"
      "github.com/minio/minio-go/v7"
      "github.com/minio/minio-go/v7/pkg/credentials"
  )
  ```
- Use Go 1.21+ for compatibility and security.

## Connection Management

- **Basic Client Initialization**: Create a client with endpoint, credentials, and SSL settings. Use `context.Background()` for non-cancellable operations.
  ```go
  ctx := context.Background()
  endpoint := "play.min.io"  // Or your MinIO server
  accessKeyID := "Q3AM3UQ867SPQQA43P2F"  // From Viper or env
  secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
  useSSL := true

  minioClient, err := minio.New(endpoint, &minio.Options{
      Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
      Secure: useSSL,
  })
  if err != nil {
      log.Fatalf("Failed to create MinIO client: %v", err)
  }
  defer minioClient.Close()
  ```
  - **Endpoint**: Include port if non-standard (e.g., "localhost:9000").
  - **Credentials**: Use `credentials.NewStaticV4` for static; `credentials.NewIAM` for dynamic (e.g., AWS IAM roles).
  - **SSL**: Set `Secure: true` for production; use custom `Transport` for TLS configs.
- **Ping for Connectivity**: Verify the connection with `FGetObject` or a simple `StatObject` on a known bucket.
  ```go
  if err := minioClient.StatObject(ctx, "testbucket", "testfile", minio.StatObjectOptions{}); err != nil {
      log.Fatalf("MinIO ping failed: %v", err)
  }
  ```
- **Connection Pooling**: The SDK handles internal pooling; configure via `http.Client` transport if needed.
  ```go
  transport := http.DefaultTransport.(*http.Transport).Clone()
  transport.MaxIdleConns = 100
  minioClient, err := minio.New(endpoint, &minio.Options{
      Creds:  creds,
      Secure: useSSL,
      Transport: transport,
  })
  ```
- **Viper Integration**: Load credentials and endpoint from config.
  ```go
  endpoint := viper.GetString("minio.endpoint")
  accessKeyID := viper.GetString("minio.access_key")
  secretAccessKey := viper.GetString("minio.secret_key")
  useSSL := viper.GetBool("minio.ssl")
  ```
  - **Config File (config.yaml)**:
    ```yaml
    minio:
      endpoint: play.min.io
      access_key: Q3AM3UQ867SPQQA43P2F
      secret_key: zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG
      ssl: true
    ```

## Project Structure

Organize MinIO-related code for modularity and testability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with MinIO initialization
  ├── internal/
  │   ├── storage/             # MinIO client setup and operations
  │   │   ├── minio.go         # Client wrapper and repository logic
  │   │   └── models.go        # Object metadata structs
  │   ├── handlers/            # Gin handlers for upload/download
  │   └── services/            # Business logic integrating storage
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define MinIO interactions in a `storage` package using a repository pattern.
- Use interfaces for storage access to enable mocking:
  ```go
  type ObjectStorage interface {
      UploadObject(ctx context.Context, bucket, object string, data []byte) error
      DownloadObject(ctx context.Context, bucket, object string) ([]byte, error)
  }
  ```

## Operations Best Practices

- **Bucket Management**: Create buckets with location (e.g., "us-east-1"); check existence first.
  ```go
  bucketName := "mybucket"
  location := "us-east-1"
  exists, err := minioClient.BucketExists(ctx, bucketName)
  if err != nil {
      return fmt.Errorf("bucket check failed: %w", err)
  }
  if !exists {
      if err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
          return fmt.Errorf("bucket creation failed: %w", err)
      }
  }
  ```
- **Upload Objects**: Use `PutObject` for small files; `PutObjectWithContext` for context support. Stream large files with `PutObject` and `io.Reader`.
  ```go
  objectName := "testfile"
  data := []byte("Hello, MinIO!")
  _, err := minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{
      ContentType: "application/octet-stream",
  })
  if err != nil {
      return fmt.Errorf("upload failed: %w", err)
  }
  ```
  - **Multipart Upload**: For >5MB files, use `NewMultipartUpload` and `UploadPart` for resumability.
  ```go
  // Multipart upload example
  upload, err := minioClient.StartMultipartUpload(ctx, bucketName, objectName, minio.PutObjectOptions{})
  for i, part := range parts {  // parts is []io.ReadSeeker
      uploadInfo, err := minioClient.UploadPart(ctx, bucketName, objectName, upload.UploadID, i+1, part, int64(part.Len()), minio.PutObjectOptions{})
      // Handle uploadInfo
  }
  _, err = minioClient.CompleteMultipartUpload(ctx, bucketName, objectName, upload, parts, minio.PutObjectOptions{})
  ```
- **Download Objects**: Use `GetObject` with `io.ReadCloser`; stream to avoid memory bloat.
  ```go
  object, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
  if err != nil {
      return nil, fmt.Errorf("download failed: %w", err)
  }
  defer object.Close()
  data, err := io.ReadAll(object)
  ```
- **List Objects**: Use `ListObjects` with prefixes and recursion for efficient pagination.
  ```go
  for object := range minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
      Prefix:     "prefix/",
      Recursive:  true,
  }) {
      // Process object.Key, object.Size
  }
  ```
- **Delete Objects**: Use `RemoveObject` for single; `RemoveObjects` for bulk.
  ```go
  err := minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
  ```
- **Presigned URLs**: Generate temporary access URLs for uploads/downloads.
  ```go
  url, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, 24*time.Hour, nil)
  // Use url.String() for client-side upload
  ```

## Security

- **Credentials**: Never hardcode; use Viper/env or `credentials.NewEnvCredentials()`. Rotate regularly.
  ```go
  creds := credentials.NewEnvCredentials()
  minioClient, err := minio.New(endpoint, &minio.Options{Creds: creds})
  ```
- **TLS/SSL**: Always enable `Secure: true`; provide custom certs via `Transport`.
  ```go
  tlsConfig := &tls.Config{RootCAs: caCertPool}
  transport := &http.Transport{TLSClientConfig: tlsConfig}
  minioClient, err := minio.New(endpoint, &minio.Options{
      Creds:     creds,
      Secure:    true,
      Transport: transport,
  })
  ```
- **IAM Policies**: Use MinIO's policy-based access; attach to users/groups.
- **Encryption**: Enable server-side encryption with `PutObjectOptions.ServerSideEncryption`.
  ```go
  opts := minio.PutObjectOptions{
      ServerSideEncryption: minio.NewSSEKMS([]byte("my-kms-key"), nil),
  }
  ```
- **Bucket Policies**: Set via MinIO Console or SDK for fine-grained access.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Repository Pattern**:
  ```go
  package storage

  type MinioRepo struct {
      Client *minio.Client
  }

  func (r *MinioRepo) UploadObject(ctx context.Context, bucket, object string, data []byte) error {
      _, err := r.Client.PutObject(ctx, bucket, object, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{
          ContentType: "application/octet-stream",
      })
      return err
  }

  // In handlers
  func UploadHandler(repo *MinioRepo) gin.HandlerFunc {
      return func(c *gin.Context) {
          file, _ := c.FormFile("file")
          f, _ := file.Open()
          defer f.Close()
          if err := repo.UploadObject(c.Request.Context(), "mybucket", file.Filename, f); err != nil {
              c.JSON(500, gin.H{"error": err.Error()})
              return
          }
          c.JSON(200, gin.H{"message": "uploaded"})
      }
  }
  ```
- **Viper Configuration**: As shown; gitignore secrets.
- **Cobra CLI**: Initialize client in root command.
  ```go
  var client *minio.Client

  var rootCmd = &cobra.Command{
      PersistentPreRun: func(cmd *cobra.Command, args []string) {
          creds := credentials.NewStaticV4(viper.GetString("minio.access_key"), viper.GetString("minio.secret_key"), "")
          var err error
          client, err = minio.New(viper.GetString("minio.endpoint"), &minio.Options{Creds: creds, Secure: true})
          if err != nil {
              log.Fatalf("MinIO client failed: %v", err)
          }
      },
  }
  ```
- **JWT Authentication**: Store signed URLs in JWT claims for secure access.
- **Zap Logging**: Log operations.
  ```go
  logger.Info("Object uploaded", zap.String("bucket", bucket), zap.String("object", object))
  if err != nil {
      logger.Error("Upload failed", zap.Error(err))
  }
  ```
- **Prometheus Metrics**: Track uploads/downloads.
  ```go
  var (
      uploadDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
          Name:    "app_minio_upload_duration_seconds",
          Help:    "Duration of MinIO uploads",
          Buckets: prometheus.DefBuckets,
      }, []string{"bucket"})
  )

  func (r *MinioRepo) UploadObject(ctx context.Context, bucket, object string, data []byte) error {
      start := time.Now()
      defer func() { uploadDuration.WithLabelValues(bucket).Observe(time.Since(start).Seconds()) }()
      // Upload logic
  }
  ```

## Testing

- **Unit Tests**: Mock `minio.Client` with `uber-go/mock` or use `minio.NewTestClient` for in-memory testing.
  ```go
  import "github.com/minio/minio-go/v7/pkg/minio"

  func TestUpload(t *testing.T) {
      // Use minio.NewTestClient or mock
      mockClient := mocks.NewMockClient(ctrl)
      mockClient.EXPECT().PutObject(gomock.Any(), "bucket", "obj", gomock.Any(), gomock.Any(), gomock.Any()).Return(&minio.UploadInfo{}, nil)
      repo := &MinioRepo{Client: mockClient}
      err := repo.UploadObject(context.Background(), "bucket", "obj", []byte("data"))
      assert.NoError(t, err)
  }
  ```
- **Integration Tests**: Use Dockerized MinIO via `testcontainers-go`.
  ```go
  import "github.com/testcontainers/testcontainers-go/modules/minio"

  func TestIntegration(t *testing.T) {
      ctx := context.Background()
      container, err := minio.Run(ctx, "minio/minio:latest")
      if err != nil {
          t.Fatalf("Failed to start MinIO: %v", err)
      }
      defer container.Terminate(ctx)
      endpoint, _ := container.Endpoint(ctx, "9000")
      creds := credentials.NewStaticV4("minioadmin", "minioadmin", "")
      client, _ := minio.New(endpoint, &minio.Options{Creds: creds, Secure: false})
      repo := &MinioRepo{Client: client}
      // Run tests
  }
  ```

## Performance Tips

- **Multipart for Large Files**: Use for >5MB to avoid timeouts; set `MultipartSize` based on network.
- **Streaming**: Always stream uploads/downloads with `io.Reader`/`io.Writer` to handle large objects.
- **Presigned URLs**: Offload uploads to clients for scalability.
- **Bucket Regions**: Set regions for multi-region setups.
- **Error Handling**: Distinguish transient (retry with `avast/retry-go`) vs permanent errors.

## Additional Resources

- Official MinIO Go SDK GitHub: https://github.com/minio/minio-go
- Go Client API Reference: https://pkg.go.dev/github.com/minio/minio-go/v7
- Quickstart Guide: https://min.io/docs/minio/linux/developers/go/minio-go.html
- Security Best Practices: https://blog.min.io/s3-security-access-control/

Adhere to these practices to ensure secure and efficient MinIO integration with the Go SDK. Update this document as new features or best practices emerge.
```