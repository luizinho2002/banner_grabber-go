# 🚩 Banner Grabber (Service Detector)

A lightweight reconnaissance tool built in **Go** to identify the software and version running on a specific network port.

## 🛠️ Key Concepts
* **TCP Connection:** Established using ``net.DialTimeout` for efficiency.
* **Service Banners:** Captures the initial identification string senty by services (SSH, FTP, etc.).
* **Safety Deadlines:** Uses `SetReadDeadline` to prevent the tool from hanging on silent services.

## 💻 How to Run
1. Confiugure the `targetIP` and `targetPort` in `main.go`.
2. Run the tool:
  ```bash
     go run main.go
  ```

# 📊 Sample Out
`[*] Banner from part 22: SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13`

