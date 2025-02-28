# Use a base Python image or any other image
FROM python:3.12

# Set the working directory
WORKDIR /app

COPY . .

# sex
ENV PYTHONUNBUFFERED=1

# Copy the requirements file and install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Run the Python script
CMD ["python3", "/app/run.py"]

