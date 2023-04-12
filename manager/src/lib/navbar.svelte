<script lang="ts">
    import { onMount } from "svelte";
    import { defaults, update } from "$lib/stores/defaults";
    import type { LayoutData } from './$types';

    onMount(() => {
        maxWidth = window.innerWidth / 2;
    });

    export let data: LayoutData;
    const minWidth = 199;
    let maxWidth = minWidth;
    let selected = data.menus[0].url;

    const onMouseDown = (e: MouseEvent) => {
        e.preventDefault();
        if (e.button !== 0) return;
        window.addEventListener('mouseup', onMouseUp);
        // window.addEventListener('touchend', onMouseUp);
        window.addEventListener('mousemove', onMouseMove);
        // window.addEventListener('touchmove', onMouseMove);
    }

    const onMouseMove = (e: MouseEvent) => {
        e.preventDefault();
        if ((e as MouseEvent).button !== 0) return;
        if (e.x <= minWidth || e.x >= maxWidth) return;
        update(obj => obj.sidebarWidth = e.x);
    }

    const onMouseUp = (e: MouseEvent) => {
        e.preventDefault();
        window.removeEventListener('mousemove', onMouseMove);
        // window.removeEventListener('touchmove', onMouseMove);
        window.removeEventListener('mouseup', onMouseUp);
        // window.removeEventListener('touchend', onMouseUp);
    }
</script>

<nav class="flex-none h-full relative" style="width: {$defaults.sidebarWidth}px;">
    <div class="bg-transparent h-full w-[8px] absolute -right-[5px] cursor-col-resize"
         on:mousedown={onMouseDown}></div>
    <div class="w-full h-full p-2">
        {#each data.menus as menu}
            <div class="rounded-md h-[30px] px-2 cursor-pointer flex items-center
                 {selected !== menu.url ? 'hover:bg-skin-accent/90' : ''}"
                 class:shadow-lg={selected === menu.url}
                 class:text-white={selected === menu.url}
                 class:hover:text-gray-200={selected !== menu.url}
                 class:bg-skin-accent={selected === menu.url}>
                {menu.name}
            </div>
        {/each}
    </div>
</nav>
