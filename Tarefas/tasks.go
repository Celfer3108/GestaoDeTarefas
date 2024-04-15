package tarefas

type Task struct {
	ID          string
	Name        string   //nome da tarefa
	Status      string   //concluída, em processo ou a iniciar
	Priority    string   //alta, média ou baixa
	DueDate     string   //data final de entrega
	Description string   //descrição da tarefa
	Assignee    Assignee //usuário responsável pela execução da tarefa
}
