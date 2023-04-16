<script lang="ts">
    import Search from "../images/icons/search.svg?component";
    import { defaults, update } from "./stores/defaults";
    import type { Menus } from "./types/menus";
    import { page } from '$app/stores';
    import MenuItem from './menuItem.svelte'

    export let data: Menus;
    const minWidth = 199;
    let selected: String;
    $: selected = $page.url.pathname;

    const onMouseDown = (e: MouseEvent) => {
        e.preventDefault();
        if (e.button !== 0) return;
        window.addEventListener('mouseup', onMouseUp);
        window.addEventListener('mousemove', onMouseMove);
    }

    const onMouseMove = (e: MouseEvent) => {
        e.preventDefault();
        if ((e as MouseEvent).button !== 0) return;
        if (e.x <= minWidth || e.x >= window.innerWidth / 2) return;
        update(obj => obj.sidebarWidth = e.x);
    }

    const onMouseUp = (e: MouseEvent) => {
        e.preventDefault();
        window.removeEventListener('mousemove', onMouseMove);
        window.removeEventListener('mouseup', onMouseUp);
    }
</script>

<nav class="flex-none h-full relative" style="width: {$defaults.sidebarWidth}px;">
    <div class="bg-transparent dark:hover:bg-gray-600/10 hover:bg-gray-400/10
            h-full w-[8px] absolute -right-[5px] cursor-col-resize"
         on:mousedown={onMouseDown}></div>
    <div class="w-full h-full p-2">
        <div class="relative">
            <label for="email" class="absolute h-[32px] px-2 flex items-center">
                <Search class="dark:fill-gray-300 dark:gray-500 scale-75"/>
            </label>
            <input id="email" type="text" name="email" placeholder="Search..."
                   class="h-[32px] mb-2 bg-gray-500/30 dark:bg-gray-500/20 rounded-md px-2 pl-[34px]
                   text-black dark:text-white border-gray-600
                   focus:placeholder-gray-500/60 placeholder-gray-500"/>
        </div>

        {#each data.menus as menu}
            <MenuItem data={menu} {selected}/>
        {/each}
    </div>
</nav>
