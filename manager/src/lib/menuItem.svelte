<script lang="ts">
    import Arrow from "../images/icons/arrow.forward.fill.svg?component";
    import type { MenuItem } from "./types/menus";
    import { page } from '$app/stores';
    import { slide } from 'svelte/transition';
    export let data: MenuItem;
    export let selected: String;
    export let subItem: boolean = false;
    export let opening: boolean = false;

    function toggleChildren() {
        opening = !opening
        if (!opening) {
            selected = data.url;
        } else {
            selected = $page.url.pathname;
        }
    }
</script>

{#if data.children}
    <div class="rounded-md select-none h-[38px] px-3 cursor-pointer flex
                 items-center transition duration-100 gap-[8px] justify-between
                 {selected !== data.url ? 'hover:bg-orange-800/10 dark:hover:bg-gray-800/60 hover:font-bold'
                 : 'bg-skin-accent shadow-lg drop-shadow-lg text-white'}" on:click={toggleChildren}>
        <div class="flex gap-[8px] w-full items-center h-full">
            <img alt={data.name} class="w-[20px] h-[20px]" src={data.icon}/>
            {data.name}
        </div>
        <div class="{opening ? 'rotate-90' : 'rotate-0'} {selected ? 'fill-gray-500' : 'fill-gray-300'}
            transition-all ease-linear duration-500 scale-75">
            <Arrow/>
        </div>
    </div>
    {#if opening}
        <div transition:slide="{{ duration: 500 }}">
            {#each data.children as menu}
                <svelte:self bind:data={menu} bind:selected={selected} subItem={true}/>
            {/each}
        </div>
    {/if}
{:else}
    <a href={data.url} class="rounded-md select-none h-[38px] px-3 cursor-pointer flex items-center transition duration-100 gap-[8px]
                 {selected !== data.url ? 'hover:bg-orange-800/10 dark:hover:bg-gray-800/60 hover:font-bold'
                 : 'bg-skin-accent shadow-lg drop-shadow-lg text-white'} {subItem ? 'pl-[40px]' : ''}"
            on:click={() => {selected = data.url}}>
        <img alt={data.name} class="w-[20px]" src={data.icon}/>
        {data.name}
    </a>
{/if}
