<script lang="ts">
    export let title: string;
    export let value: string | number;
    export let description: string;
    export let valueColor: string = "text-blue-600";
</script>

<div class="p-6 bg-white rounded-lg shadow-md">
    <h2 class="text-xl font-semibold text-gray-800">{title}</h2>
    <p class="mt-2 text-3xl font-bold {valueColor}">{value}</p>
    <p class="mt-1 text-sm text-gray-500">{description}</p>
</div>
