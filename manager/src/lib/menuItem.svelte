<script lang="ts">
    import Arrow from "../images/icons/chevron.right.svg?component";
    import type { MenuItem } from "./types/menus";
    import { page } from '$app/stores';
    import { slide } from 'svelte/transition';
    import { sineInOut } from "svelte/easing";
    import { onMount } from "svelte";
    export let data: MenuItem;
    export let selected: String;
    export let subItem: boolean = false;
    export let opening: boolean = false;
    let isSelected: boolean;
    $: isSelected = selected === data.url;

    onMount(() => {
        if (!data.children) {
            return
        }

        if (data.children.some(m => m.url === selected)) {
            opening = true;
        }
    });

    function toggleChildren() {
        if (!data.children) {
            return
        }

        opening = !opening
        if (!opening && data.children?.some(m => m.url === selected)) {
            selected = data.url;
        } else {
            selected = $page.url.pathname;
        }
    }
</script>

{#if data.children}
    <div class="rounded-md select-none h-[38px] px-3 cursor-pointer flex
                 items-center transition duration-100 gap-[8px] justify-between
                 {isSelected ? 'bg-skin-accent shadow-lg drop-shadow-lg text-white' :
                 'hover:bg-gray-800/10 dark:hover:bg-gray-800/60 hover:font-bold'}"
         on:click={toggleChildren} on:keydown={() => {}}>
        <div class="flex gap-[8px] w-full items-center h-full">
            <img alt={data.name} class="w-[20px] h-[20px]" src={data.icon}/>
            {data.name}
        </div>
        <div class="{opening ? 'rotate-90' : 'rotate-0'} {isSelected ? 'fill-gray-300' : 'dark:fill-gray-500 fill-gray-400'}
            transition-all ease-in-out duration-500 scale-75">
            <Arrow/>
        </div>
    </div>
    {#if opening}
        <div transition:slide="{{ duration: 500, easing: sineInOut }}">
            {#each data.children as menu}
                <svelte:self bind:data={menu} bind:selected={selected} subItem={true}/>
            {/each}
        </div>
    {/if}
{:else}
    <a href={data.url} class="rounded-md select-none h-[38px] px-3 cursor-pointer flex items-center transition duration-100 gap-[8px]
                 {isSelected ? 'bg-skin-accent shadow-lg drop-shadow-lg text-white' :
                 'hover:bg-gray-800/10 dark:hover:bg-gray-800/60 hover:font-bold'} {subItem ? 'pl-[35px]' : ''}"
            on:click={() => {selected = data.url}}>
        <img alt={data.name} class="w-[20px] h-[20px]" src={data.icon}/>
        {data.name}
    </a>
{/if}
