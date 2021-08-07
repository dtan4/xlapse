# xlapse-image-cleaner

Clean up all downloaded images from S3 bucket except `images.yaml` which is the image list used by Distributor, and `movie.gif` created by GifMaker

## Usage

```bash
go build
./DRY_RUN=true ./xlapse-image-cleaner <BUCKET_NAME>
./xlapse-image-cleaner <BUCKET_NAME>
```
