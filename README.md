# form


URLS

Templates

| Verb | Path                        | Usage                                          |
|------|-----------------------------|------------------------------------------------|
| GET  | /templates/new              | Show the form builder                          |
| POST | /templates                  | Create a new form template                     |
| GET  | /templates                  | Return the list of form templates              |
| POST | /templates/{id}             | Update a form template                         |
| DEL  | /templates/{id}             | Delete a form template                         |

Forms

| Verb | Path                        | Usage                                          |
|------|-----------------------------|------------------------------------------------|
| GET  | /form/new?template={id}     | Display a new form based on template {id}      |
| POST | /forms                      | Save new form                                  |
| GET  | /forms                      | Return the list of my forms                    |
| POST | /forms/{id}                 | Update a form                                  |
| DEL  | /forms/{id}                 | Delete a form                                  |

Review

| Verb | Path                        | Usage                                          |
|------|-----------------------------|------------------------------------------------|
TODO

[ ] Form edit
[ ] Form display
[ ] Review module
[ ] Communications
[ ] User management