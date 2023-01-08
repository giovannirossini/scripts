import tinify
import os

tinify.key = os.getenv("TINIFY_API_KEY")

def tiny_from_file(input_file, output_file):
    source = tinify.from_file(input_file)
    source.to_file(output_file)

def tiny_from_url(url, output_file):
    source = tinify.from_url(url)
    source.to_file(output_file)

def resize_image(input_file, width, heigth, output_file):
    source = tinify.from_file(input_file)
    resized = source.resize(
        method="fit",
        width=width,
        height=heigth
    )
    resized.to_file(output_file)

def preserving_metadata(input_file, preserve, output_file):
    source = tinify.from_file(input_file)
    copyrighted = source.preserve(preserve)
    copyrighted.to_file(output_file)

def save_to_s3(input_file, bucket, output_file):
    source = tinify.from_file(input_file)
    source.store(
        service="s3",
        aws_access_key_id=os.getenv("AWS_ACCESS_KEY"),
        aws_secret_access_key=os.getenv("AWS_SECRET_KEY"),
        region=os.getenv("AWS_DEFAULT_REGION"),
        path=bucket + "/" + output_file
    )

def save_to_gcs(input_file, bucket, output_file):
    source = tinify.from_file(input_file)
    source.store(
        service="gcs",
        gcp_access_token=os.getenv("GCP_ACCESS_TOKEN"),
        path=bucket + "/" + output_file
    )

tiny_from_file("input.png","output.png")
tiny_from_url("https://example.com/image.jpeg", "image.jpeg")
resize_image("input.jpg", 200, 250, "output.jpg")
preserving_metadata("input.jpg","copyright", "output.jpg")
save_to_s3("input.png", "my-bucket", "tinify/1.png")
save_to_gcs("input.png", "my-bucket", "tinify/1.png")