# 🕵️‍♂️ CredGear

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](#)
[![License](https://img.shields.io/github/license/Pnkcaht/CredGear)](#)
[![CI](https://img.shields.io/github/actions/workflow/status/Pnkcaht/CredGear/ci.yml)](#)
[![Security Tool](https://img.shields.io/badge/category-security--tool-red)](#)

CredGear is a high-performance data sanitization tool designed to extract, normalize, and deduplicate credentials from unstructured text sources, producing clean, analysis-ready data.

---

## Table of Contents

- [What is CredGear](#what-is-credgear)
- [Key Features](#key-features)
- [Use Cases](#use-cases)
- [How It Works](#how-it-works)
- [Installation](#installation)
- [Output Format](#output-format)
- [Performance](#performance)
- [Security Notice](#security-notice)
- [Contributing](#contributing)
- [License](#license)

---

## What is CredGear

CredGear processes raw credential dumps such as logs, leaks, text files, and scraped content, converting them into structured, normalized JSON while removing noise and duplicates.

It is designed for correctness, performance, and automation.

<img width="1231" height="670" alt="CredGear example output" src="https://github.com/user-attachments/assets/bc56d2ac-c413-4567-a0ee-8269d7ec021f" />

## Key Features

- High-performance processing of large datasets
- Extraction of credentials from unstructured text
- Normalization of usernames, emails, and domains
- Efficient deduplication
- Structured JSON output
- Suitable for automation pipelines
- Written in Go for portability and speed

## Use Cases

CredGear is suitable for:

- Security research
- Incident response
- Breach and leak analysis
- Credential hygiene pipelines
- Threat intelligence preprocessing
- Automation workflows

## How It Works

1. Input ingestion of unstructured credential data  
2. Pattern-based credential extraction  
3. Normalization to canonical formats  
4. Deduplication across datasets  
5. Emission of structured JSON output  

## Installation

Binary releases and documentation are available on the official website:

**https://credgear.io**  
_(replace with the actual site URL if different)_

Output Format

CredGear outputs normalized JSON entries similar to the following:


{
  "username": "john.doe",
  "email": "john.doe@example.com",
  "password": "********",
  "source": "input_file"
}



The output is designed to integrate easily with analysis tools, SIEMs, and automation pipelines.

## Performance

CredGear is optimized for large-scale credential datasets using streaming processing and efficient memory usage.

Benchmarking and performance metrics will be published as the project evolves.

## Security Notice

This tool is intended for defensive, research, and authorized use only.

Users are responsible for ensuring compliance with applicable laws and regulations when handling sensitive data.

## Contributing

Contributions are welcome via issues and pull requests.

Please follow Go best practices and responsible security guidelines.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
