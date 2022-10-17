# Promotions-API
<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#used-technology-stack">Used Technology Stack</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation-and-usage">Installation and Usage</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Sample demonstration project that is about serving promotions data from a backend service.
The proejct built on top of the REST API architecture.
Used Go (Gin, GORM) for code. Postgres as a storage. AWS (EC2, Security Group, Load Balancer) for deployment
and handling the high load,Docker for containerization and Prometheus for metrics collection.

Live Demo
* http://ec2-3-84-214-210.compute-1.amazonaws.com:1321/promotions/1

Project also uses Go program for Bulk data inserting into Postgres.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Used Technology Stack

[![Go][Go]][Go-url]
[![AWS][AWS]][AWS-url]
[![Docker][Docker]][Docker-url]
[![Postgres][Postgres]][Postgres-url]
[![Git][Git]][Git-url]
[![Prometheus][Prometheus]][Prometheus-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## Getting Started

Follow the instructions for setting up the prject locally.

### Prerequisites

Before jumping into project cloning make sure to have 

* Go > 1.17
  ```sh
  Download from here https://go.dev/dl/
  ```
* Postgres
  ```sh
  Download from here https://www.postgresql.org/download/
  ```
* Docker
  ```sh
  Download from here https://www.docker.com/
  ```
* Make
  ```sh
  Download from here https://www.gnu.org/software/make/ (Not Mandatody)
  ```
* Curl
  ```sh
  Download from here https://www.gnu.org/software/make/ (Not Mandatody)
  ```

### Installation and Usage

Instructions for setting up and running the project

1. Clone the repo
   ```sh
   git clone https://github.com/p44dev/Promotions-API.git
   ```
2. Configure variables in .env file or export them
   ```sh
   For reference: .env.sample
   ```
3. Go to the Promotions-API/Bulk directory and run command below for bulk insert
   ```sh
   make insert
   ```
4. Go to the Promotions-API directory and run commands below for running app as a container
   (Not a mandatory option, for testing you can just run -> go run main.go)
   ```sh
   make dockerize
   make run-container
   ```
5. Fetch the data from the backend server.
   ```sh
   curl http://localhost:1321/promotions/1
   ```
6. See Prometheus Logs
   ```sh
   curl http://localhost:9091/metrics
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Roadmap

- [ ] Use AWS EC2 more poewrful Instace or deploy into Beanstalk (Now using 1 cpu from free tier)
- [ ] Add Cron Job for inserting bulk data every 30 minutes and swap tables.
- [ ] Add POST | PATCH | DELETE etc.. requests into REST service 
- [ ] Containerize bulk inserting application and postgres

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Anton - [@LinkedIn](https://www.linkedin.com/in/anton-papazyan-718512147/) - antonpapazian@gmail.com

Project Link: [https://github.com/p44dev/Promotions-API](https://github.com/p44dev/Promotions-API)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Acknowledgments

* [TechWorld with Nana](https://www.youtube.com/c/TechWorldwithNana) :rocket:
* [GIN](https://github.com/gin-gonic/gin/releases) :construction_worker:
* [Img Shields](https://shields.io) :fire:


## Thank You :blush:

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[Go]: https://img.shields.io/badge/Go-000000?style=for-the-badge&logo=GO&logoColor=blue
[Go-url]: https://go.dev/
[AWS]: https://img.shields.io/badge/AWS-000000?style=for-the-badge&logo=amazon&logoColor=orange
[AWS-url]: https://aws.amazon.com/
[Docker]: https://img.shields.io/badge/Docker-000000?style=for-the-badge&logo=docker&logoColor=blue
[Docker-url]: https://www.docker.com/
[Postgres]: https://img.shields.io/badge/Postgres-000000?style=for-the-badge&logo=postgresql&logoColor=blue
[Postgres-url]: https://www.postgresql.org
[Git]: https://img.shields.io/badge/Git-000000?style=for-the-badge&logo=git&logoColor=red
[Git-url]: https://git-scm.com/
[Prometheus]: https://img.shields.io/badge/prometheus-000000?style=for-the-badge&logo=prometheus&logoColor=red
[Prometheus-url]: https://prometheus.io/
