var ul = document.createElement("ul");
add_li("Сделать задание #3 по web-программированию");
var div = document.getElementById("root");
div.appendChild(ul);

function delete_from_tree()
{
    this.parentNode.remove();
}

var input = document.createElement("input");
input.id = "add_task_input";
var button_for_task = document.createElement("button");
button_for_task.id = "add_task";
button_for_task.innerHTML = "Добавить";
button_for_task.addEventListener("click", function () {
    add_li(input.value);
});

div.appendChild(input);
div.appendChild(button_for_task);

function add_li(text)
{
    var li = document.createElement("li");
    var span = document.createElement("span");
    span.innerHTML = text;
    li.appendChild(span);
    ul.appendChild(li);
    var button_for_delete = document.createElement("button");
    button_for_delete.innerHTML = "Удалить";
    button_for_delete.addEventListener("click", delete_from_tree);
    li.appendChild(button_for_delete);
}


