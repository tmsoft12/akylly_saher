<script>
  import { writable } from "svelte/store";
  import { onMount } from "svelte";

  // Route store
  export const route = writable("dashboard");

  // Import pages
  import Dashboard from "./pages/Dashboard.svelte";
  import Settings from "./pages/Settings.svelte";
  import Users from "./pages/Users.svelte";
  import History from "./pages/History.svelte";

  // Define routes
  const routes = {
    dashboard: Dashboard,
    settings: Settings,
    users: Users,
    history: History,
  };

  // Sync route with URL
  function updateRouteFromPath() {
    const path = window.location.pathname.replace("/", "") || "dashboard";
    route.set(routes[path] ? path : "dashboard"); // Fallback to "dashboard" if path is invalid
  }

  onMount(() => {
    updateRouteFromPath();
    window.addEventListener("popstate", updateRouteFromPath);

    return () => {
      window.removeEventListener("popstate", updateRouteFromPath);
    };
  });

  function navigateTo(path) {
    window.history.pushState({}, "", `/${path}`);
    route.set(path);
  }
</script>

<svelte:head>
  <link
    href="https://fonts.googleapis.com/icon?family=Material+Icons"
    rel="stylesheet"
  />
</svelte:head>

<div class="flex h-screen">
  <!-- Sidebar -->
  <div class="bg-gray-800 text-white w-64 p-4 flex flex-col justify-between">
    <div class="space-y-4">
      <h2 class="text-xl font-bold">Awtopark</h2>
      <nav class="space-y-2">
        <button
          on:click={() => navigateTo("dashboard")}
          class="flex items-center p-2 rounded hover:bg-gray-700 w-full text-left space-x-2"
        >
          <span class="material-icons">dashboard</span>
          <span>Dashboard</span>
        </button>
        <button
          on:click={() => navigateTo("users")}
          class="flex items-center p-2 rounded hover:bg-gray-700 w-full text-left space-x-2"
        >
          <span class="material-icons">person</span>
          <span>Profile</span>
        </button>
        <button
          on:click={() => navigateTo("settings")}
          class="flex items-center p-2 rounded hover:bg-gray-700 w-full text-left space-x-2"
        >
          <span class="material-icons">settings</span>
          <span>Settings</span>
        </button>
        <button
          on:click={() => navigateTo("history")}
          class="flex items-center p-2 rounded hover:bg-gray-700 w-full text-left space-x-2"
        >
          <span class="material-icons">schedule</span>
          <span>Geçmiş</span>
        </button>
      </nav>
    </div>

    <!-- Footer -->
    <p class="text-sm text-center text-gray-500 mt-4">
      © 2024 ETUT Awtopark Dolandyryşy
    </p>
  </div>

  <!-- Main Content -->
  <div class="flex-1 bg-gray-100 p-4">
    <!-- Render the dynamic component -->
    <svelte:component this={routes[$route] || Dashboard} />
  </div>
</div>
