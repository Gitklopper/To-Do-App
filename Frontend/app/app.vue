<template>
  <div>
    <input v-model="newTodoTitle" placeholder="Neues Todo"/>
    <button @click="addTodo">Hinzufügen</button>
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        {{ todo.title }} - {{ todo.progression }}
      </li>
    </ul>
  </div>
</template>

<script setup>

const todos = ref([])
const newTodoTitle = ref("")

onMounted(async () => {
  const result = await $fetch('http://localhost:8080/todo')
  todos.value = result || []
})

async function addTodo() {
  if (!newTodoTitle.value.trim()) return
  const task = await $fetch('http://localhost:8080/todo', {
    method: 'POST',
    body: { title: newTodoTitle.value }
  })
  todos.value.push(task)
  newTodoTitle.value = ""
}

</script>
