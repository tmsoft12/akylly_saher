<script lang="ts">
  import { onMount } from "svelte";
  import InputSearch from "../components/InputSearch.svelte";
  interface User {
    id: number;
    name: string;
    plate: string;
    phone: string;
    status: string;
  }

  let users: User[] = [];
  let isLoading = true;
  let error: string | null = null;

  // API'den veri çekme
  onMount(async () => {
    try {
      const response = await fetch("https://api.example.com/users"); // Buraya kendi API endpoint'inizi ekleyin
      if (!response.ok) {
        throw new Error("API isteği başarısız oldu");
      }
      users = await response.json();
    } catch (err) {
      error = (err as Error).message;
    } finally {
      isLoading = false;
    }
  });
</script>

<div class="p-6 space-y-6">
  <header class="flex items-center justify-between mb-8">
    <div>
      <h1 class="text-3xl font-bold text-gray-800">Şahsylar</h1>
      <p class="mt-2 text-gray-600">Ähli hasaba alnan ulanyjylar</p>
    </div>
    <button
      class="px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700"
    >
      <span class="material-icons align-middle mr-1">add</span>
      Täze Ulanyjy
    </button>
  </header>
  <div class="p-6 bg-white rounded-lg shadow-md">
    <div class="flex justify-between mb-4">
      <InputSearch />
      <div class="flex space-x-2">
        <button
          class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200"
        >
          <span class="material-icons align-middle">filter_list</span>
        </button>
        <button
          class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200"
        >
          <span class="material-icons align-middle">download</span>
        </button>
      </div>
    </div>

    <table class="w-full">
      <thead>
        <tr class="text-left border-b">
          <th class="px-4 py-3 text-gray-600">Ady</th>
          <th class="px-4 py-3 text-gray-600">Ulag Belgisi</th>
          <th class="px-4 py-3 text-gray-600">Telefon</th>
          <th class="px-4 py-3 text-gray-600">Ýagdaýy</th>
          <th class="px-4 py-3 text-gray-600">Amallar</th>
        </tr>
      </thead>
      <tbody>
        {#each users as user}
          <tr class="border-b">
            <td class="px-4 py-3">{user.name}</td>
            <td class="px-4 py-3">
              <div class="plate">
                <span class="plate-format">{user.plate}</span>
              </div>
            </td>
            <td class="px-4 py-3">{user.phone}</td>
            <td class="px-4 py-3">
              <span
                class="px-2 py-1 text-sm rounded-full {user.status === 'active'
                  ? 'bg-green-100 text-green-800'
                  : 'bg-gray-100 text-gray-800'}"
              >
                {user.status === "active" ? "Işjeň" : "Işjeň däl"}
              </span>
            </td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="p-1 text-blue-600 hover:text-blue-800">
                  <span class="material-icons">edit</span>
                </button>
                <button class="p-1 text-red-600 hover:text-red-800">
                  <span class="material-icons">delete</span>
                </button>
              </div>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>

    <div class="flex items-center justify-between mt-4">
      <p class="text-sm text-gray-600">Jemi: {users.length} ulanyjy</p>
      <div class="flex space-x-2">
        <button
          class="px-3 py-1 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200"
        >
          Öňki
        </button>
        <button
          class="px-3 py-1 text-white bg-blue-600 rounded-md hover:bg-blue-700"
        >
          1
        </button>
        <button
          class="px-3 py-1 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200"
        >
          Indiki
        </button>
      </div>
    </div>
  </div>
</div>
