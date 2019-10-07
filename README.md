# PLANS API

# O que é ?

Bem eu acabei me empolgando um pouco, e acabei fazendo "over-engineering" e criando meio que um "cms mutante", que exibe tanto o portal da vivo quanto a claro, porém cada um chama o dado da respectiva api variando somente o css do mesmo, através de uma classe de tema (lembra o wordpress) que é injetada de componente para componente através da arvere (que acarreta em um acomplamento desnecessário), no entanto garantiu uma rapida implementação de layout com reaproveitamento do css.

Não utilizei uma biblioteca de estado já que achei desnecessária pelo tamanho do projeto, optei por utilizar o própio setstate.

# Para inicializar o servidor local

1. Utilize o script docker para instanciar o servidor em mongodb

2. Crie um novo `.env` usando o `.env.example` de base.

3. Execute `sh reset-server.sh` na pasta principal, ele irá :

   1. Entrar na pasta `$ ./migrations` e executar `go run main.go down && go run main.go up` assim resetando o banco de dados com as migrations de preços iniciais.

   2. Voltar para pasta server e inicializar o servidor `go run main.go`.

4. Utilize-a

# Para inicializar Front End

1. Entre na pasta front_end e execute `yarn dev` ele irá instalar os pacotes e iniciar o front-end em `localhost:3000`.

# Pontos a melhorar.

- Adicionar aviso para navegadoresnão modernos.
- Dividir handlers dentre handlers e controllers para separar as regras dentre eles.
- Desacoplar os dois temas em front ends separadas porém utilizando um repositório com os componentes em comum para serém reutilizados para assim conseguir utilizar um cdn e reduzir tempo de carregamento do site.
- UX. Adicionar shimmer no carregamento.
- Adicionar mais casos nos testes unitários.
- Escrever testes de handlers utilizando [o guia](https://blog.questionable.services/article/testing-http-handlers-go/).
- Rotas para atualizar recursos da api.

# Vulnerabilidades.

- XSS (Cross Site Scripting) na página de detalhes possui uma vulnerabilidade ao injetar html no content de detalhes.

# Considerações finais

Eu ainda não tinha feito um hello world com golang, e só criei na vida um hello world com reactjs não fazia ideia do quão simples as duas tecnologias.

GOLANG é incrível mto simples estruturar o código e a forma com que ele trata erros é bem interessante por não se ter um `try catch` o esquema de verificar `err` em cada function e passando para cima casa bem com a sintaxe dele, uma ótima escolha para microserviços.

Eu sei que fiz besteira criando esse monstrinho ehehehe.
