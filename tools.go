package main

//go:generate sh -c "mkdir -p worker && /bin/echo -n $'package worker\n\nimport \"fmt\"\n\nfunc Run() {\n\tfmt.Println(\"Running\")\n}\n' > worker/worker.go"

