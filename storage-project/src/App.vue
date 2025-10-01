<template>
  <div class="warehouse-inventory">
    <h1>Складской учет 📦</h1>

    <div class="add-item-form">
      <h2>Добавить товар</h2>
      <div class="input-group">
        <input
          v-model="newItem.name"
          type="text"
          placeholder="Название товара (например, 'Кабель USB-C')"
          @keyup.enter="addItem"
        />
        <input
          v-model.number="newItem.quantity"
          type="number"
          min="1"
          placeholder="Количество"
          @keyup.enter="addItem"
        />
        <button @click="addItem" :disabled="!isFormValid">
          Добавить
        </button>
      </div>
      <p v-if="error" class="error-message">{{ error }}</p>
    </div>

    <hr />

    <div class="inventory-table">
      <h2>Текущий Запас (Всего: {{ totalItems }} ед.)</h2>
      
      <p v-if="inventory.length === 0" class="empty-message">
        Склад пуст! Добавьте первую позицию.
      </p>

      <table v-else>
        <thead>
          <tr>
            <th>ID</th>
            <th>Название Товара</th>
            <th>Количество</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in inventory" :key="item.ID">
            <td>{{ item.ID }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.quantity }}</td>
            <td>
              <button @click="removeItem(item.ID)" class="remove-btn">
                Удалить
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios'; // <-- Импорт Axios
const API_URL = "http://localhost:8080/api/v1/inventory/"

const error = ref('');
// Имитация данных из БД
const inventory = ref([]);
// Реактивное состояние для формы
const newItem = ref({
  name: '',
  quantity: 1,
});

// Вычисляемое свойство для проверки валидности формы
const isFormValid = computed(() => {
  return newItem.value.name.trim() !== '' && newItem.value.quantity > 0;
});

// Вычисляемое свойство для общего количества всех единиц товаров
const totalItems = computed(() => {
  return inventory.value.reduce((sum, item) => sum + item.quantity, 0);
});

const fetchInventory = async () => {
  error.value = '';
  try {
    // GET-запрос: загружаем все товары
    const response = await axios.get(API_URL);
    inventory.value = response.data;
  } catch (err) {
    console.error("Ошибка загрузки склада:", err);
    error.value = 'Не удалось загрузить данные склада. Проверьте, запущен ли бэкенд на порту 8080.';
  }
};

// Метод добавления товара
const addItem = async () => {
  error.value = '';

  if (!isFormValid.value) {
    error.value = 'Заполните название и укажите количество (больше 0).';
    return;
  }

  // Данные для отправки на бэкенд (только name и quantity)
  const dataToSend = {
    name: newItem.value.name.trim(),
    quantity: newItem.value.quantity,
  };

  try {
    // POST-запрос: создаем новый товар
    await axios.post(API_URL, dataToSend);

    // Перезагружаем список, чтобы увидеть новый товар, включая его ID
    await fetchInventory(); 
    
    // Очистка формы
    newItem.value.name = '';
    newItem.value.quantity = 1;
    
  } catch (err) {
    console.error("Ошибка при добавлении товара:", err);
    // Выводим ошибку с бэкенда (например, конфликт имени)
    const backendError = err.response?.data?.error || 'Неизвестная ошибка при добавлении.';
    error.value = backendError;
  }
};

// Метод удаления товара
const removeItem = async (id) => {
  error.value = '';
  try {
    // DELETE-запрос: удаляем товар по ID
    await axios.delete(`${API_URL}${id}`);
    
    // Перезагружаем список
    await fetchInventory();
    
  } catch (err) {
    console.error("Ошибка при удалении товара:", err);
    error.value = 'Не удалось удалить товар. Возможно, он уже был удален.';
  }
};

onMounted(() => {
  fetchInventory();
})
</script>

<style scoped>
.warehouse-inventory {
  max-width: 900px;
  margin: 20px auto;
  padding: 30px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  border: 1px solid #e0e0e0;
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  background-color: #ffffff;
}

h1 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 25px;
}

hr {
  margin: 25px 0;
  border: 0;
  border-top: 1px dashed #cccccc;
}

/* --- Форма добавления --- */
.add-item-form {
  padding: 20px;
  border: 1px solid #dcdcdc;
  border-radius: 6px;
  background-color: #f9f9f9;
}

.input-group {
  display: flex;
  gap: 10px;
  align-items: center;
}

.add-item-form input[type="text"] {
  flex-grow: 2;
  padding: 10px;
}

.add-item-form input[type="number"] {
  flex-grow: 1;
  max-width: 100px;
  padding: 10px;
  text-align: center;
}

.add-item-form input {
  border: 1px solid #ccc;
  border-radius: 4px;
}

.add-item-form button {
  padding: 10px 20px;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
  white-space: nowrap; /* Чтобы кнопка не переносилась */
}

.add-item-form button:hover:not(:disabled) {
  background-color: #369c67;
}

.add-item-form button:disabled {
  background-color: #a0a0a0;
  cursor: not-allowed;
}

.error-message {
  color: #e74c3c;
  margin-top: 10px;
  font-weight: 500;
}

/* --- Таблица товаров --- */
.inventory-table h2 {
    margin-bottom: 15px;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
  background-color: #ffffff;
}

th, td {
  border: 1px solid #dddddd;
  padding: 12px 15px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
  font-weight: bold;
  color: #333;
}

tr:nth-child(even) {
  background-color: #f9f9f9;
}

.remove-btn {
  padding: 6px 12px;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-size: 14px;
}

.remove-btn:hover {
  background-color: #c0392b;
}

.empty-message {
  text-align: center;
  color: #7f8c8d;
  font-style: italic;
  padding: 20px;
  border: 1px dashed #ccc;
  border-radius: 4px;
  background-color: #fcfcfc;
}
</style>