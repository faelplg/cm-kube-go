Curso de Microsserviços:
# Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços

## Módulo: DEVOPS

### Exercício: Utilizando K8s
* Autor: Rafael Goulart
  * E-mail cadastrado: rafaelgoulart@residuall.org
  * E-mail pessoal: faelplg@gmail.com

#### Tarefa 1: Servidor Web - Nginx
* Nome do artefato: `cm-kube-nginx`
* [Repositório do GitHub](https://github.com/faelplg/cm-kube-nginx)

Resolução:
* Criação do `deployment.yaml`;
  * Configuração simples especificando um container com `nginx:1.18-alpine`;
* Criação do `configmap.yaml`;
  * Neste arquivo quero definir o conteúdo do arquivo html que irá exibir a mensagem "Code.Education Rocks!".
* Criação de um volume no `deployment.yaml` que utiliza o ConfigMap criado;
* Montagem desse volume na pasta `/usr/share/nginx/html`, copiando a alteração realizada;
* Criação de um serviço para expor o NGINX;

#### Tarefa 2: Configuração do MySQL
* Nome do artefato: `cm-kube-mysql`
* [Repositório do GitHub](https://github.com/faelplg/cm-kube-mysql)

Resolução:
* Criação do `deployment.yaml`;
  * Configuração simples especificando um container com `mysql:5.7`;
* Criação do `persistent-volume.yaml` para especificar um PersistentVolumeClaim, no qual será definido um volume persistente para o DB;
* Criação do volume persistente dentro do `deployment.yaml`;
* Montagem desse volume na pasta `/var/lib/mysql`;
* Atribuição de um secret por linha de comando (`kube-mysql-secret`) para esconder password do DB;
* Criação do serviço para expor o MySQL;

#### Tarefa 3: Desafio Go!
* Nome do artefato: `cm-kube-go`
* [Repositório do GitHub (este)](https://github.com/faelplg/cm-kube-go)

Resolução:
* Criação do programa em GO;
* Criação do teste do programa;
* Criação do `Dockerfile` e `cloudbuild.yaml` para processo de CI;
  * A imagem da aplicação é criada com um _build_ do programa copiado para uma imagem `scratch`;
* Publicação da imagem no Docker Hub;
* Criação de um `deployment.yaml` com a imagem publicada;
* Criação de um serviço expondo a aplicação para acesso via browser;
