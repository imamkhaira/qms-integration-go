# QMS Integration in Go

This CLI software enables the integration of China QMS Machine to our reporting server.

# Sekilas Info

## List of Machines

| Machine Serial No   | Site Name            |
| ------------------- | -------------------- |
| QMSK-V02-B002-N0001 | IWK Shah Alam        |
| QMSK-V02-B002-N0002 | IWK KL Pantai        |
| QMSK-V02-B002-N0003 | IWK KL (UTC)         |
| QMSK-V02-B002-N0004 | IWK Hulu Langat      |
| QMSK-V02-B002-N0005 | IWK N.Sembilan       |
| QMSK-V02-B002-N0006 | IWK Melaka (UTC)     |
| QMSK-V02-B002-N0007 | IWK Kluang           |
| QMSK-V02-B002-N0008 | IWK Skudai           |
| QMSK-V02-B002-N0009 | IWK Skudai (UTC)     |
| QMSK-V02-B002-N0010 | IWK Pahang (UTC)     |
| QMSK-V02-B002-N0011 | IWK Alor Setar (UTC) |
| QMSK-V02-B002-N0012 | IWK Sg Petani        |
| QMSK-V02-B002-N0013 | IWK Klang            |
| QMSK-V02-B002-N0014 | IWK Seberang Prai    |
| QMSK-V02-B002-N0015 | IWK Ipoh (UTC)       |
| QMSK-V02-B002-N0016 | IWK Manjung          |
| QMSK-V02-B002-N0017 | IWK Taiping          |
| QMSK-V02-B002-N0018 | IWK Terengganu       |
| QMSK-V02-B002-N0019 | IWK Kelantan         |

## How it Works

1. This software runs as a Windows scheduled task.
2. It will query the database for records of the previous day,
3. then uploads it to the reporting server.
4. if the connection failed, it will attempt again at the next run cycle.

## Limitations

1. cannot receive data from server to the QMS machine.
2. does not know whether the ticket is auto-printed or not.
3. banyak issue maaaa

## Future Roadmap

1. add feature to receive data from the server
2. bugfix maaa

# Development

## Tools & Setup

1. Microsoft SQL Server 2008 r2
2. Golang >= 1.20.0
3. msodbc = 17, CANNOT USE v18 AND ABOVE!
4. Windows 10 to build

## Credentials in the QMS machine

- host: `localhost`
- username: `sa`
- password: `sa123456789@`

## Setting Up Development

1. setup sql server machine
2. update the kiosk.ini file as above
3. run `go mod install`

To run the program once: `go run cmd/qms-cli/main.go`
To run the build process: `go run scripts/build.go`

# DISCLAIMER

I wrote this program as part of my job.
If this program is used by parties other than myself, then I am not responsible for any problems it may cause, be it bricked devices, server crash, corrupt database, nuclear war, wrath of God, or you getting fired because Boss thinks it is your fault. Please read the source code carefully, and make sure you understand what it does before using it! YOU are choosing to run OR make modifications to this, and if you blame me in any way for what happens to you, I will laugh at you.
