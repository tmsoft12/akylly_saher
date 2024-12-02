<script lang="ts">
  let settings = {
    hourlyRate: 10,
    maxDuration: 24,
    currency: "TMT",
    notifications: true,
    autoBackup: true,
  };
</script>

<div class="p-6 space-y-6">
  <header class="mb-8">
    <h1 class="text-3xl font-bold text-gray-800">Sazlamalar</h1>
    <p class="mt-2 text-gray-600">Ulgamyň sazlamalaryny dolandyryň</p>
  </header>

  <div class="grid gap-6 md:grid-cols-2">
    <div class="p-6 bg-white rounded-lg shadow-md">
      <h2 class="mb-4 text-xl font-semibold">Töleg Sazlamalary</h2>
      <div class="space-y-4">
        <div>
          <label
            for="hourlyRate"
            class="block mb-2 text-sm font-medium text-gray-700"
          >
            Sagatlaýyn Töleg
          </label>
          <input
            id="hourlyRate"
            type="number"
            bind:value={settings.hourlyRate}
            class="w-full px-3 py-2 border rounded-md"
          />
        </div>
        <div>
          <label
            for="currency"
            class="block mb-2 text-sm font-medium text-gray-700"
          >
            Pul Birligi
          </label>
          <select
            id="currency"
            bind:value={settings.currency}
            class="w-full px-3 py-2 border rounded-md"
          >
            <option value="TMT">TMT</option>
            <option value="USD">USD</option>
            <option value="EUR">EUR</option>
          </select>
        </div>
      </div>
    </div>

    <div class="p-6 bg-white rounded-lg shadow-md">
      <h2 class="mb-4 text-xl font-semibold">Ulgam Sazlamalary</h2>
      <div class="space-y-4">
        <div class="flex items-center justify-between">
          <span class="text-sm font-medium text-gray-700">Duýduryşlar</span>
          <label
            for="notifications"
            class="relative inline-flex items-center cursor-pointer"
          >
            <input
              id="notifications"
              type="checkbox"
              bind:checked={settings.notifications}
              class="sr-only peer"
            />
            <div
              class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"
            ></div>
          </label>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm font-medium text-gray-700"
            >Awtomatik ätiýaçlama</span
          >
          <label
            for="autoBackup"
            class="relative inline-flex items-center cursor-pointer"
          >
            <input
              id="autoBackup"
              type="checkbox"
              bind:checked={settings.autoBackup}
              class="sr-only peer"
            />
            <div
              class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"
            ></div>
          </label>
        </div>
      </div>
    </div>
  </div>

  <div class="flex justify-end mt-6">
    <button
      class="px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700"
    >
      Ýatda Sakla
    </button>
  </div>
</div>
