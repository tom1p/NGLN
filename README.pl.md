# NGLN - Next Generation Logistic Network

Polski | [English](README.md)

## Opis projektu
NGLN to system CRM nowej generacji, zaprojektowany w celu uproszczenia zarządzania sieciami logistycznymi. Projekt wykorzystuje architekturę mikroserwisową, zapewniając skalowalność, modularność i łatwość utrzymania.

---

## ✅ Status projektu
### Funkcje
- [ ] **API Gateway**: Centralne routowanie i uwierzytelnianie. *(W trakcie)*
- [ ] **Frontend**: Interfejs użytkownika. *(Zaplanowane)*
- [ ] **User Service**: Zarządzanie użytkownikami i rolami. *(Zaplanowane)*
- [ ] **Sales Service**: Obsługa klientów, leadów i sprzedaży. *(Zaplanowane)*

---

## 🛠️ Technologie
- **Backend**: Golang, Laravel, Nest.js
- **Frontend**: Next.js
- **Bazy danych**: MySQL, PostgreSQL
- **Orkiestracja**: Kubernetes
- **Konteneryzacja**: Docker

---

## 📁 Struktura projektu
```plaintext
ngln/
│
├── api-gateway/                # API Gateway (Golang)
├── user-service/               # Serwis zarządzania użytkownikami (Laravel)
├── sales-service/              # Serwis sprzedaży i CRM (Nest.js)
├── frontend/                   # Interfejs użytkownika (Next.js)
├── deployment/                 # Pliki wdrożeniowe Kubernetes
└── README.md                   # Dokumentacja
```

---

## 🚀 Instalacja
### Wymagania
- Docker >= 20.10
- Kubernetes (Minikube, Rancher Desktop lub klaster w chmurze)
- Helm >= 3.0
- Node.js >= 20.0
- Composer >= 2.0
- Go >= 1.20

### Kroki
1. Sklonuj repozytorium:
   ```bash
   git clone https://github.com/YourUsername/ngln.git
   cd ngln
   ```
2. Wdrożenie usług w Kubernetes:
   - Postępuj zgodnie z instrukcjami w katalogach poszczególnych usług (`api-gateway/`, `user-service/` itp.).

---

## 🧪 Testowanie
1. Upewnij się, że usługi są wdrożone i działają w klastrze Kubernetes.
2. Sprawdź logi w celu debugowania:
   ```bash
   kubectl logs <pod-name> -n ngln
   ```

---

## 📜 Licencja
Projekt jest licencjonowany na podstawie Apache License 2.0. Szczegóły znajdziesz w pliku [LICENSE](LICENSE).

---

## 📋 Plan

### **Milestone 1: MVP**

- [ ] Implementacja systemu logowania i rejestracji użytkowników.
- [ ] Zaimplementowanie API dla CRUD użytkowników.
- [ ] Stworzenie prostego frontendu.

### **Milestone 2: Backend i Frontend**

- [ ] Rozbudowa backendu o dodatkowe funkcje.
- [ ] Dalszy rozwój frontendu i zapewnienie responsywności.
- [ ] Przygotowanie środowiska na Kubernetes.

### **Milestone 3: CI/CD (Integracja i Deployment)**

- [ ] Skonfigurowanie pipeline’u CI/CD, automatyczne testy i wdrożenia na środowisko staging.

### **Milestone 4: Rozwój UX/UI**

- [ ] Rozwój UX/UI.
- [ ] Ulepszenie interfejsu użytkownika.

### **Milestone 5: Testy i optymalizacja**

- [ ] Testowanie API, frontend, integracja z Kubernetes i Docker.
- [ ] Optymalizacja aplikacji pod kątem wydajności.

### **Milestone 6: Integracja z wybranym systemem płatności (przyszły plan - opcjonalnie)**

- [ ] Integracja z wybranym systemem płatności.
- [ ] Wdrożenie i przetestowanie podstawowych płatności.

### **Milestone 7: Finalizacja**

- [ ] Skonfigurowanie pipeline’u CI/CD, automatyczne testy i wdrożenia na środowisko produkcyjne.
- [ ] Wdrożenie aplikacji na produkcję.

