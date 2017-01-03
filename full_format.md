## PESSOA
### Endpoints
| Method      | URL           	| Descrição
| ---         | ---           	| ---
| `GET`       | **/people**    	| Lista todas as pessoas
| `POST`      | **/people**    	| Cadastra pessoa
| `PUT`       | **/people/:id** | Altera pessoa pelo ID
| `GET`       | **/people/:id** | Busca pessoa pelo ID
| `DELETE`    | **/people/:id** | Exclui pessoa pelo ID

### Objeto
| Nome                  	| Tipo        	| Descrição
| ---                   	| ---         	| ---
| **id**                    | `uuid`      	| ID identificador da pessoa
| **code**                  | `int`       	| Código identificador da pessoa
| **name**                  | `string`    	| Nome da pessoa
| **city_name**             | `string`    	| Nome da cidade
| **company_name**          | `string`    	| Razão social
| **address**               | `string`    	| Endereço
| **number**                | `string`    	| Número do endereço
| **complement**            | `string`    	| Complemento do endereço
| **district**              | `string`    	| Bairro
| **zip**                   | `string`    	| CEP
| **birth_date**            | `timestamp` 	| Data de nascimento
| **cpf**                   | `string`    	| CPF
| **rg**                    | `string`    	| RG
| **gender**                | `char`      	| Sexo
| **business_phone**        | `string`    	| Telefone comercial
| **home_phone**            | `string`    	| Telefone residencial
| **mobile_phone**          | `string`    	| Telefone celular
| **cnpj**                  | `string`    	| CNPJ
| **state_inscription**     | `string`    	| Inscrição estadual
| **phone**                 | `string`    	| Telefone
| **fax**                   | `string`    	| Fax
| **email**                 | `string`    	| E-mail
| **website**               | `string`    	| Site
| **observations**          | `string`    	| Observações
| **registered_at**         | `timestamp` 	| Data de registro
| **registered_by**         | `uuid`      	| Quem registrou
| **type**                  | `char`      	| Tipo de pessoa

---

## TAREFA
### Endpoints
| Method      | URL           	| Descrição
| ---         | ---           	| ---
| `GET`       | **/tasks**    	| Lista todas as tarefas
| `POST`      | **/tasks**    	| Cadastra tarefa
| `PUT`       | **/tasks/:id**  | Altera tarefa pelo ID
| `GET`       | **/tasks/:id**  | Busca tarefa pelo ID
| `DELETE`    | **/tasks/:id**  | Exclui tarefa pelo ID

### OBJETO
| Nome                  	| Tipo        	    | Descrição
| ---                   	| ---         	    | ---
| **id**                    | `uuid`      	    | ID identificador da tarefa
| **code** 				    | `int` 		    | Código identificador da tarefa
| **title**                 | `string` 		    | Título
| **due**                   | `time.Time`  	    | Data de vencimento
| **visualized**            | `bool` 		    | Visualizada
| **completed_at**          | `*time.Time` 	    | Data de conclusão
| **registered_at**         | `time.Time` 	    | Data de cadastro
| **category**              | `TaskCategory`    | Categoria
| **registered_by**         | `Person`	        | Cadastrado por
| **person**                | `Person`	        | Pessoa
| **assignee**              | `Person`	        | Responsável
| **task_historics**        | `[]TaskHistorics` | Histórico

---

## CATEGORIA DE TAREFA
### Endpoints
| Method      | URL           	            | Descrição
| ---         | ---                 	    | ---
| `GET`       | **/task_categories**    	| Lista todas as categorias
| `POST`      | **/task_categories**    	| Cadastra categoria
| `PUT`       | **/task_categories/:id**    | Altera categoria pelo ID
| `GET`       | **/task_categories/:id**    | Busca categoria pelo ID
| `DELETE`    | **/task_categories/:id**    | Exclui categoria pelo ID

### OBJETO
| Nome              | Tipo      | Descrição
| ---               | ---       | ---
| **id**            | `uuid`    | ID identificador da categoria
| **description**   | `string`  | Descrição da categoria

---

## HISTÓRICO DE TAREFA
### OBJETO
| Nome              	| Tipo      	| Descrição
| ---               	| ---       	| ---
| **id**            	| `uuid`    	| ID identificador da categoria
| **comment**			| `string` 		| Comentário
| **registered_at**		| `time.Time`	| Data de cadastro
| **registered_by**		| `Person`		| Cadastrado por