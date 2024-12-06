
# NGLN - Next Generation Logistic Network

Polski | [English](README.md)

## Opis projektu
NGLN to system CRM nowej generacji, zaprojektowany w celu uproszczenia zarządzania sieciami logistycznymi. Projekt wykorzystuje architekturę mikroserwisową, zapewniając skalowalność, modularność i łatwość utrzymania.

---

## ✅ Status projektu
### Funkcje
- [ ] **API Gateway**: Centralne routowanie i uwierzytelnianie. *(W trakcie)*
- [ ] **User Service**: Zarządzanie użytkownikami i rolami. *(Zaplanowane)*
- [ ] **Sales Service**: Obsługa klientów, leadów i sprzedaży. *(Zaplanowane)*
- [ ] **Frontend**: Interfejs użytkownika w Next.js. *(Zaplanowane)*

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

## 📋 Przyszłe cele
- [ ] Wdrożenie kontroli dostępu opartej na rolach (RBAC) w User Service.
- [ ] Integracja obsługi płatności w Sales Service.
- [ ] Dodanie zaawansowanych modułów analitycznych i raportów.
