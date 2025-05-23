# Zeiterfassung CLI

Ein einfaches Kommandozeilen-Tool zur Zeiterfassung und zum Auslesen der aktuellen Zeit von einer Webseite.

## Features

- **Start**: Beginnt eine Zeiterfassungssitzung.
- **Stop**: Beendet die aktuelle Zeiterfassungssitzung.
- **Status**: Zeigt den aktuellen Status der Zeiterfassung.
- **Fetch**: Liest die aktuelle Zeit von einer angegebenen Webseite aus.

## Installation

1. Repository klonen:
   ```sh
   git clone https://github.com/dein-benutzername/zeiterfassung.git
   cd zeiterfassung
   ```

2. Abhängigkeiten installieren:
   ```sh
   go mod tidy
   ```

3. Build:
   ```sh
   go build -o zeiterfassung
   ```

## Nutzung

```sh
./zeiterfassung <befehl> [argumente]
```

### Beispiele

- **Zeit von Webseite auslesen:**
  ```sh
  ./zeiterfassung fetch https://example.com
  ```

- **Zeiterfassung starten:**
  ```sh
  ./zeiterfassung start
  ```

- **Zeiterfassung stoppen:**
  ```sh
  ./zeiterfassung stop
  ```

- **Status anzeigen:**
  ```sh
  ./zeiterfassung status
  ```

## Flags

Jeder Befehl unterstützt ggf. weitere Flags. Mit `--help` erhältst du eine Übersicht:
```sh
./zeiterfassung <befehl> --help
```

## Mitwirken

Pull Requests und Vorschläge sind willkommen!

## Lizenz

MIT License