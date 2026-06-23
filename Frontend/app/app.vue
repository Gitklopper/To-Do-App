<template>
  <div>
    <input v-model="newTodoTitle" placeholder="Neues Todo"/>
    <button @click="addTodo">Hinzufügen</button>
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        <span> {{ todo.title }} </span>
        <select :value="todo.progression" @change="updateStatus(todo.id, $event.target.value)">
          <option value="open">Offen</option>
          <option value="in_progress">In Bearbeitung</option>
          <option value="done">Fertig</option>
        </select>
        <select :value="todo.priority" @change="updatePriority(todo.id, $event.target.value)">
          <option value="neither">Keine Priorität</option>
          <option value="urgent">Dringend</option>
          <option value="important">Wichtig</option>
          <option value="urgent_important">Dringend & Wichtig</option>
        </select>
        <span> Erstellt: {{ formatDateTime(todo.created_at) }} </span>
        <span> |  Geändert: {{ formatDateTime(todo.updated_at) }} </span>
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

async function updateStatus(id, status) {
  const result = await $fetch('http://localhost:8080/todo/update', {
    method: 'PATCH',
    body: { id, status }
  })
  todos.value = result
}

async function updatePriority(id, priority) {
  const result = await $fetch('http://localhost:8080/todo/update', {
    method: 'PATCH',
    body: { id, priority }
  })
  todos.value = result
}

function formatDateTime(dateStr) {
  const date = new Date(dateStr)
  const time = date.toLocaleTimeString('de-DE', { hour: '2-digit', minute: '2-digit' })
  const day = date.toLocaleDateString('de-DE')
  return `${time} ${day}`
}

</script>
