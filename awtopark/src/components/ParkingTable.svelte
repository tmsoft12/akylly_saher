<script lang="ts">
  import { onMount } from "svelte";
  import InputSearch from "./InputSearch.svelte";

  // ParkingEntry tipini doğru şekilde tanımlayın
  interface ParkingEntry {
    plate: string;
    entryTime: string;
    exitTime: string;
  }

  // parkingData dizisinin doğru tipini belirtin
  let parkingData: ParkingEntry[] = [];
  let isLoading = true;
  let error: string | null = null;
  let searchTerm = ""; // Arama terimi

  onMount(() => {
    let socket: WebSocket;

    const connect = () => {
      socket = new WebSocket("ws://192.168.0.104:3000/ws/plate");

      socket.onopen = () => {
        console.log("WebSocket bağlantısı başarılı.");
        isLoading = false;
        error = null;
      };

      socket.onmessage = (event) => {
        const data: ParkingEntry[] = JSON.parse(event.data);
        if (Array.isArray(data)) {
          parkingData = data;
        } else {
          error = "Geçersiz veri formatı alındı.";
        }
      };

      socket.onerror = () => {
        error = "WebSocket bağlantısında bir hata oluştu.";
      };

      socket.onclose = () => {
        console.log("WebSocket bağlantısı kapandı. Yeniden bağlanıyor...");
        setTimeout(connect, 3000); // 3 saniye sonra yeniden bağlan
      };
    };

    connect();

    return () => {
      socket.close();
    };
  });
</script>

<div class="p-6 bg-white rounded-lg shadow-md border overflow-hidden h-full">
  <div class="flex justify-between">
    <h2 class="text-xl font-semibold text-gray-800">Giriş Maglumatlary</h2>
  </div>

  {#if isLoading}
    <p>Loading...</p>
  {:else if error}
    <p class="text-red-600">Hata: {error}</p>
  {:else if parkingData.length === 0}
    <p class="text-gray-500">Maglumat yok.</p>
  {:else}
    <div
      class="overflow-y-scroll w-full h-full scrollbar-thin scrollbar-thumb-gray-400 scrollbar-track-gray-200"
    >
      <table class="w-full mt-4 text-left table-auto">
        <thead>
          <tr class="border-b">
            <th class="px-4 py-2 text-gray-600">Ulag Belgisi</th>
            <th class="px-4 py-2 text-gray-600">Giren Wagty</th>
            <th class="px-4 py-2 text-gray-600">Çykan Wagty</th>
          </tr>
        </thead>
        <tbody>
          {#each parkingData.filter((entry) => entry.plate
              .toLowerCase()
              .includes(searchTerm.toLowerCase())) as entry}
            <tr class="border-b">
              <td class="px-4 py-2 text-gray-800">
                <div class="plate">
                  <span class="plate-format">{entry.plate}</span>
                </div>
              </td>
              <td class="px-4 py-2">
                <div class="time-box time-box-entry">
                  <span class="material-icons">arrow_downward</span>
                  <span class="time">{entry.entryTime}</span>
                </div>
              </td>

              <td class="px-4 py-2">
                <div class="time-box time-box-exit">
                  <span class="material-icons">arrow_upward</span>
                  {#if entry.exitTime === ""}
                    <span class="time">Içerde</span>
                  {:else}
                    <span class="time">{entry.exitTime}</span>
                  {/if}
                </div>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>
