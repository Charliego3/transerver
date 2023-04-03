<div class="relative bg-skin-fill overflow-hidden h-full">
    <div class="bg-image-light dark:bg-image-dark absolute inset-0 h-full w-full object-cover opacity-30"></div>
    <div class="absolute inset-0 bg-gradient-to-br from-skin-background via-skin-background to-transparent"></div>
    <div class="relative h-full">
        <div class="flex flex-row pt-5 pl-10 font-bold text-md gap-4">
            <h1 class="font-splash text-2xl">Transerver</h1>
            <h1 class="font-splash text-2xl">Manager</h1>
        </div>

        <div class="grid grid-cols-1 gap-4 px-10 h-full content-center">
            <slot/>
        </div>
    </div>
</div>
