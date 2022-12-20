![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)

<div align="center" >
  <h2>CRUD Books</h2>
</div>
<h4 align="center">Uma API que fornece um CRUD para livros!</h4>
<p align="center">
 <a href="#sobre">Sobre</a> •
 <a href="#funcionalidades">Funcionalidades</a> • 
 <a href="#resultado">Resultado</a> • 
 <a href="#autor">Autor</a>
</p>

<h2 id="sobre">🧾 Sobre</h2>
<p>Este projeto foi desenvolvido em Golang como forma de estudar e aprender mais, na prática, sobre a linguagem! 👨‍💻</p>
<p>Essa é uma aplicação de um CRUD de livros, onde é possível realizar todas as operações básicas, como: criar um livro, pesquisar por um livro, editar um livro e até mesmo apagar um livro. Também podemos receber todos os livros em um array.</p>

<h2 id="funcionalidades">⚙Funcionalidades</h2>
<ul>
  <li>Criar um livro</li>
  <li>Pesquisar por todos os livros</li>
  <li>Pesquisar pelos livros (pelo id)</li>
  <li>Editar um livro</li>
  <li>Apagar um livro</li>
</ul>

<h2 id="resultado">📱 Resultado</h2>
<h2>Criar um livro</h2>
<h3>Request</h3>


```bash
curl --request POST \
  --url http://localhost:1337/books \
  --header 'Content-Type: application/json' \
  --data '{
  "title": "Book Title",
  "author": "Book Author"
}'
```
![POST_CREATE](https://user-images.githubusercontent.com/80720221/208697758-280934e8-e64d-45cc-9701-5fb0ae5f14b0.gif)


<h2>Buscar por todos os livros</h2>
<h3>Request</h3>


```bash
curl --request GET \
  --url http://localhost:1337/books
```

![GET_ALL](https://user-images.githubusercontent.com/80720221/208700588-d26bf135-3ae3-444a-8a87-141470fd4c8c.gif)

<h2>Buscar por um livro baseado no id</h2>
<h3>Request</h3>

```bash
curl --request GET \
  --url http://localhost:1337/books/{id}
```

![GET_ID](https://user-images.githubusercontent.com/80720221/208700886-90dbab93-72f6-4d55-a6bf-4f92f5a79c38.gif)

<h2>Editar um livro</h2>
<h3>Request</h3>

```bash
curl --request PUT \
  --url http://localhost:1337/books/{id} \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Book Title",
	"author": "Book Author"
}'
```

![PUT](https://user-images.githubusercontent.com/80720221/208701911-6fd928aa-5d69-4981-b4ba-ef1e9acfc129.gif)

<h2>Excluir um livro</h2>
<h3>Request</h3>

```bash
curl --request DELETE \
  --url http://localhost:1337/books/{id}
```

![DELETE](https://user-images.githubusercontent.com/80720221/208702530-ba1e2b57-5e20-4281-8057-6dde08204a37.gif)


<h2 id="autor">👨‍💻 Autor</h2>
<table>
  <tr>
    <td align="center">
      <a href="https://github.com/davidsonmarra">
        <img src="https://github.com/davidsonmarra.png?size=100" width="100px;" alt="Davidson Marra"/><br>
        <sub>
          <b>Davidson Marra</b>
        </sub>
      </a>
    </td>
  </tr>
</table>
