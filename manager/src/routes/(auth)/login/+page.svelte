<script lang="ts">
    import Logo from "/src/images/icons/logo.svg?component";

    import { enhance } from '$app/forms';
    import type { SubmitFunction } from './$types';
    import { invalidateAll } from "$app/navigation";

    const loginHandler: SubmitFunction = ({ from, data, action, cancel }) => {
        invalidateAll();
    }

    let codeFocused = false;
    function toggleCodeFocus() {
        codeFocused = !codeFocused;
    }
</script>

<div class="w-full flex flex-col justify-center items-center gap-4 sm:gap-1 pb-10 -mt-32">
    <Logo class="w-16 fill-skin-fill"/>
    <h2 class="mt-6 text-center text-4xl font-bold text-gray-900 dark:text-gray-300">Sign in to your account</h2>
    <p class="mt-2 text-center text-sm text-gray-600 max-w">
        Not registered yet?
        <a href="/register" class="text-sm underline decoration-sky-500 font-medium
                decoration-solid decoration-2 underline-offset-4 text-blue-600">Sign up</a>
    </p>
</div>

<form action="?/login" method="POST"
      class="flex flex-col gap-4 h-full w-full bg-gray-400 px-10
    py-8 rounded-lg shadow-lg drop-shadow-lg bg-opacity-20 dark:bg-opacity-10" use:enhance={loginHandler}>
    <div>
        <label for="uname" class="text-skin-second font-medium pl-0.5">User name</label>
        <input id="uname" type="text" name="uname" placeholder="Transerver"/>
    </div>
    <div>
        <label for="password" class="text-skin-second font-medium pl-0.5">Password</label>
        <input id="password" type="password" name="password" placeholder="enter password"/>
    </div>
    <div class="flex flex-col">
        <label for="code" class="text-skin-second font-medium pl-0.5">Email code</label>
        <div id="code" class="inline-flex rounded-lg" class:ring-1={codeFocused} class:ring-skin-accent={codeFocused}>
            <input type="text" name="code" placeholder="123456" class="rounded-r-none border-r-0 focus:ring-skin-accent" on:focusin={toggleCodeFocus} on:focusout={toggleCodeFocus}/>
            <button on:click={(e) => {e.preventDefault()}} class="bg-skin-accent text-white px-3 border-l-0 rounded-r-lg font-medium shadow-lg drop-shadow-sm focus:ring-skin-accent">Send</button>
        </div>
    </div>
    <div>
        <label for="google" class="text-skin-second font-medium pl-0.5">Google code</label>
        <input id="google" type="text" name="google" placeholder="enter google"/>
    </div>
    <div class="flex items-center text-skin-base mb-5">
        <input id="remember-me" name="remember" type="checkbox"/>
        <label for="remember-me" class="w-full ml-2 text-sm flex justify-between">
            <span class="text-skin-second">Remember me</span>
            <a href="?/forgotPassword" class="text-skin-accent font-semibold hover:underline">Forgot your password?</a>
        </label>
    </div>
    <button type="submit" class="bg-skin-accent font-bold rounded-lg py-2 text-white
            shadow-lg drop-shadow-sm hover:bg-opacity-90">Sign in
    </button>
</form>
