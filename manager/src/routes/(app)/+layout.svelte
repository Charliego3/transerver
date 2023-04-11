<script lang="ts">
    import Header from '$lib/header.svelte';
    import Navbar from '$lib/navbar.svelte';
    import type { LayoutData } from './$types';
    export let data: LayoutData;
</script>

<div class="flex flex-col h-screen divide-y dark:divide-slate-400/20 divide-gray-400/50">
    <Header/>
    <div class="h-full">
        <div class="flex h-full w-full divide-x dark:divide-slate-400/20 divide-gray-400/50">
            <Navbar {data}/>
            <div class="w-full p-2">
                <slot/>
            </div>
        </div>
    </div>
</div>
