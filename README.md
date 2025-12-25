## Жобаның архитектуралық диаграммасы

```mermaid
graph TD
    User((Пайдаланушы)) -->|HTTP/Next.js| Frontend[Frontend: Next.js]
    Frontend -->|REST API| Backend[Backend: Go API]
    Backend -->|SQL| DB[(Database: PostgreSQL)]
    
    subgraph Infrastructure
        Terraform[Terraform: IaC] -->|Басқару| K8s[Kubernetes Cluster]
        Docker[Docker: Containers] --> K8s
    end
    
    subgraph CI_CD
        Git[GitHub Push] --> Actions[GitHub Actions]
        Actions -->|Build & Test| Docker
        Actions -->|Deploy| K8s
    end
    
    Backend -.->|Мониторинг| Grafana[Grafana / Sentry]
