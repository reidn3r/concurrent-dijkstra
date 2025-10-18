# 🧩 Dijkstra Concorrente em Go

## 🚀 Problema Abordado

Encontrar o **caminho de menor custo** em um grafo grande pode ser custoso.
Quando o grafo possui muitos nós e arestas, algoritmos sequenciais são ineficientes

Essa solução aborda a idéia de explorar múltiplos nós concorrentemente, buscando **redução do tempo de execução** em grafos grandes.

---

## 💡 Implementação

* **Grafo gerado aleatoriamente** com `n` nós e `m` arestas por nó.
* **Dijkstra concorrente**: múltiplas goroutines exploram vizinhos em paralelo, com controle de threads para evitar overhead.
* Estruturas críticas (`visited` e `costSoFar`) usam **`sync.Map`**, versão thread-safe do `map`.

---

## 🖨️ Saída:

Ao executar o programa, imprime:

* Se o **caminho foi encontrado** ou não
* O **custo total do caminho** (soma dos pesos das arestas)
* O **tempo total de execução**

Exemplo:

```
Caminho encontrado! Custo total: 42
Tempo total de execução: 12.345ms
```

---

## ⚙️ Como Rodar

1. Clone o projeto e entre na pasta:

```bash
git clone <repo-url>
cd astar-concurrent
```

2. Inicialize o módulo Go e organize dependências:

```bash
go mod tidy
```

3. Execute o programa:

```bash
go run main.go
```

## 🧠 Observações
* Controle de **threads (`maxThreads`)** evita excesso de goroutines, equilibrando concorrência e desempenho.
