<script lang="ts">
    import Header from '$lib/header.svelte';
    import Sidebar from '$lib/sidebar.svelte';
    import type { LayoutData } from './$types';
    export let data: LayoutData;
</script>

<div class="flex flex-col h-screen divide-y divide-skin-fill dark:bg-skin-fill">
    <Header/>
    <div class="h-full">
        <div class="flex h-full w-full divide-x divide-skin-fill">
            <Sidebar {data}/>
            <div class="h-full w-full p-2">
                <slot/>
            </div>
        </div>
    </div>
</div>
