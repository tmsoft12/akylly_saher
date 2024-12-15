<script lang="ts">
    import { onMount } from "svelte";
    import StatCard from "../components/StatCard.svelte";
    import ParkingTable from "../components/ParkingTable.svelte";

    interface ParkingCars {
        empty_exit_count: string;
        total_car_count: string;
        totalPayments: string;
    }

    let parkingCars: ParkingCars = {
        empty_exit_count: "0",
        total_car_count: "0",
        totalPayments: "0",
    };

    let bos: number = 0;
    let isLoading = true;
    let error: string | null = null;

    let previousData: ParkingCars = {
        empty_exit_count: "0",
        total_car_count: "0",
        totalPayments: "0",
    };

    // Bildirim sesini yükle
    let notificationSound = new Audio("/sounds/not.wav");

    // Ses dosyasının düzgün yüklendiğini kontrol et
    notificationSound.oncanplaythrough = () => {
        console.log("Ses dosyası başarıyla yüklendi.");
    };
    notificationSound.onerror = (err) => {
        console.error("Ses dosyasını yüklerken bir hata oluştu:", err);
    };

    // Verilerdeki değişiklikleri kontrol etmek için fonksiyon
    const checkForChanges = (newData: ParkingCars) => {
        // Değişiklikleri kontrol et
        if (
            newData.empty_exit_count !== previousData.empty_exit_count ||
            newData.total_car_count !== previousData.total_car_count ||
            newData.totalPayments !== previousData.totalPayments
        ) {
            console.log("Veriler değişti, ses çalınacak!");
            notificationSound.play(); // Sesli bildirim
            previousData = { ...newData }; // Yeni veriyi kaydet
        }
    };

    onMount(() => {
        const ws = new WebSocket("ws://192.168.0.103:3000/ws/dashboard");

        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log("Gelen veri:", data); // Gelen veriyi konsola yazdır

            if (data.error) {
                error = data.error;
            } else {
                parkingCars.empty_exit_count = data.empty_exit_count.toString();
                parkingCars.total_car_count = data.total_car_count.toString();
                parkingCars.totalPayments = data.totalPayments.toFixed(2);

                bos = 100 - Number(parkingCars.empty_exit_count);

                // Değişiklik olup olmadığını kontrol et ve bildirim göster
                checkForChanges(parkingCars);
            }
        };

        ws.onerror = (err) => {
            console.error("WebSocket error:", err);
            error = "WebSocket bağlantısı başarısız.";
        };

        ws.onclose = () => {
            console.log("WebSocket bağlantısı kapandı.");
        };

        isLoading = false;

        return () => {
            ws.close(); // Bileşen unmount olduğunda bağlantıyı kapat.
        };
    });
</script>

<div class="p-6 space-y-6">
    <header class="flex items-center justify-between">
        <h1 class="text-3xl font-bold text-gray-800">Dashboard</h1>
    </header>

    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
        <StatCard
            title="Gunluk Hasap"
            value={parkingCars.total_car_count}
            description="Bir günde girip çykanlar"
            valueColor="text-blue-600"
        />
        <StatCard
            title="Hasap"
            value="{parkingCars.totalPayments} TMT"
            description="Bu aýyň jemindäki girdeji"
            valueColor="text-green-600"
        />
        <StatCard
            title="Ulag"
            value={parkingCars.empty_exit_count}
            description="Park ýagdaýyndaky ulaglar"
            valueColor="text-yellow-600"
        />
        <StatCard
            title="Boş Yer"
            value={bos}
            description="Awto parkdaky boş orunlar"
            valueColor="text-red-600"
        />
    </div>
    <div class="h-96">
        <ParkingTable />
    </div>
</div>
