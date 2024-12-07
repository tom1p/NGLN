# NGLN - Next Generation Logistic Network

English | [Polski](README.pl.md)

## Project Description
NGLN is a next-generation CRM system designed to simplify logistic network management. The project is built using a microservices architecture to ensure scalability, modularity, and maintainability.

---

## ✅ Project Status
### Features
- [ ] **API Gateway**: Centralized routing and authentication. *(In Progress)*
- [ ] **Frontend**: User interface. *(Planned)*
- [ ] **User Service**: User and role management. *(Planned)*
- [ ] **Sales Service**: Customer, lead, and sales handling. *(Planned)*


---

## 🛠️ Technologies
- **Backend**: Golang, Laravel, Nest.js
- **Frontend**: Next.js
- **Databases**: MySQL, PostgreSQL
- **Orchestration**: Kubernetes
- **Containerization**: Docker

---

## 📁 Project Structure
```plaintext
ngln/
│
├── api-gateway/                # API Gateway (Golang)
├── user-service/               # User management service (Laravel)
├── sales-service/              # Sales and CRM service (Nest.js)
├── frontend/                   # Frontend interface (Next.js)
├── deployment/                 # Kubernetes deployment files
└── README.md                   # Documentation
```

---

## 🚀 Installation
### Requirements
- Docker >= 20.10
- Kubernetes (Minikube, Rancher Desktop, or cloud-based cluster)
- Helm >= 3.0
- Node.js >= 20.0
- Composer >= 2.0
- Go >= 1.20

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/YourUsername/ngln.git
   cd ngln
   ```
2. Deploy services using Kubernetes:
   - Follow instructions in each service directory (`api-gateway/`, `user-service/`, etc.).

---

## 🧪 Testing
1. Ensure services are deployed and running in your Kubernetes cluster.
2. Check logs for debugging:
   ```bash
   kubectl logs <pod-name> -n ngln
   ```

---

## 📜 License
This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

---

## 📋 Plan

### **Milestone 1: MVP**

- [ ] Implement user registration and login system.
- [ ] Implement the API for user CRUD operations.
- [ ] Create a simple frontend.

### **Milestone 2: Backend and Frontend**

- [ ] Extend the backend with additional features.
- [ ] Continue developing the frontend and ensure responsiveness.
- [ ] Prepare the environment for Kubernetes.

### **Milestone 3: CI/CD (Integration and Deployment)**

- [ ] Set up CI/CD pipeline, automated tests, and deployments to the staging environment.

### **Milestone 4: UX/UI Development**

- [ ] Develop UX/UI.
- [ ] Improve the user interface.

### **Milestone 5: Testing and Optimization**

- [ ] Test the API, frontend, and integration with Kubernetes and Docker.
- [ ] Optimize the application for performance.

### **Milestone 6: Integration with Selected Payment System (future plan - optional)**

- [ ] Integrate with the selected payment system.
- [ ] Implement and test basic payments.

### **Milestone 7: Finalization**

- [ ] Set up CI/CD pipeline, automated tests, and deployments to the production environment.
- [ ] Deploy the application to production.

