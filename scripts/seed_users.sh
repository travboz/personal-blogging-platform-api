#!/bin/bash

# Define the API endpoint
url="http://localhost:7666/articles"

# Define the articles to insert
articles=(
  "{\"content\": \"The future of remote collaboration tools is being shaped by AI-driven features that improve team productivity.\", \"tags\": [\"technology\", \"remote-work\"]}"
  "{\"content\": \"Eating more plant-based foods can significantly reduce the risk of chronic diseases.\", \"tags\": [\"health\", \"sustainability\"]}"
  "{\"content\": \"Understanding compound interest is essential for long-term financial planning and wealth building.\", \"tags\": [\"finance\", \"education\"]}"
  "{\"content\": \"New startups in renewable energy are revolutionizing how we generate and store power.\", \"tags\": [\"startups\", \"sustainability\"]}"
  "{\"content\": \"The rise of ransomware attacks has made cybersecurity a top priority for businesses.\", \"tags\": [\"cybersecurity\", \"technology\"]}"
  "{\"content\": \"Online courses are making education more accessible to remote communities.\", \"tags\": [\"education\", \"remote-work\"]}"
  "{\"content\": \"Balancing screen time and physical activity is crucial for children's health in a digital age.\", \"tags\": [\"health\", \"education\"]}"
  "{\"content\": \"Fintech startups are simplifying how people manage their money with intuitive mobile apps.\", \"tags\": [\"startups\", \"finance\"]}"
  "{\"content\": \"Implementing strong password policies helps organizations protect against data breaches.\", \"tags\": [\"cybersecurity\"]}"
  "{\"content\": \"The shift to remote work is reducing urban congestion and improving work-life balance.\", \"tags\": [\"remote-work\", \"sustainability\"]}"
)


# Loop through each user and insert them using curl
for article in "${articles[@]}"; do
  curl -X POST "$url" \
    -H "Content-Type: application/json" \
    -d "$article"
  # echo -e "\Article inserted: $article\n"
done
