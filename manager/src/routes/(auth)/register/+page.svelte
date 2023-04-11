<script lang="ts">
    import Logo from "/src/images/icons/logo.svg?component";

    import { enhance } from '$app/forms';
    import type { ActionData, PageData, SubmitFunction } from './$types';

    export let data: PageData;
    export let form: ActionData;

    const registerHandler: SubmitFunction = ({ from, data, action, cancel }) => {
        const {uname, email} = Object.fromEntries(data);
        if (!uname || !email) {
            cancel();
        }
    }

    let codeFocused = false;
    function toggleCodeFocus() {
        codeFocused = !codeFocused;
    }
</script>

<div class="w-full flex flex-col justify-center items-center gap-4 sm:gap-1 pb-10 -mt-32">
    <Logo class="w-16 fill-skin-fill"/>
    <h2 class="mt-6 text-center text-4xl font-bold text-gray-900 dark:text-gray-300 px-10">Create your account</h2>
    <p class="mt-2 text-center text-sm text-gray-600 max-w">
        Already registered?
        <a href="/login" class="text-sm underline decoration-sky-500
                decoration-solid decoration-2 underline-offset-4 text-blue-600 font-medium">Sign in</a>
    </p>
</div>

{#if form?.success}
    <p>Successfully logged in! Welcome back, {data.uame}</p>
{/if}

<form action="?/register" method="POST"
      class="flex flex-col gap-3 h-full w-full bg-gray-400 px-10
    py-8 rounded-lg shadow-lg drop-shadow-lg bg-opacity-20 dark:bg-opacity-10" use:enhance={registerHandler}>
    <div>
        <label for="uname" class="text-skin-second font-medium pl-0.5">User name</label>
        <input id="uname" type="text" name="uname" placeholder="Transerver"/>
    </div>
    <div>
        <label for="email" class="text-skin-second font-medium pl-0.5">Email address</label>
        <input id="email" type="text" name="email" placeholder="transerver@mail.com"/>
    </div>
    <div class="flex flex-col">
        <label for="code" class="text-skin-second font-medium pl-0.5">Email code</label>
        <div id="code" class="inline-flex rounded-lg" class:ring-1={codeFocused} class:ring-skin-accent={codeFocused}>
            <input type="text" name="code" placeholder="123456" class="rounded-r-none border-r-0 focus:ring-skin-accent" on:focusin={toggleCodeFocus} on:focusout={toggleCodeFocus}/>
            <button on:click={() => {}} class="bg-skin-accent text-white px-3 border-l-0 rounded-r-lg font-medium shadow-lg drop-shadow-sm focus:ring-skin-accent">Send</button>
        </div>
    </div>
    <div>
        <label for="password" class="text-skin-second font-medium pl-0.5">Password</label>
        <input id="password" type="password" name="password" placeholder="enter password"/>
    </div>
    <div>
        <label for="confirm" class="text-skin-second font-medium pl-0.5">Confirm Password</label>
        <input id="confirm" type="password" name="confirm" placeholder="confirm password"/>
    </div>

    <div class="flex items-center text-skin-base">
        <input id="terms-and-privacy" name="privacy" type="checkbox"/>
        <label for="terms-and-privacy" class="ml-2 block text-sm text-skin-second">
            I agree to the
            <a href="#" class="text-skin-accent font-semibold hover:underline">Terms</a>
            and
            <a href="#" class="text-skin-accent font-semibold hover:underline">Privacy Policy</a>.
        </label>
    </div>
    <button type="submit" class="mt-5 bg-skin-accent font-bold rounded-lg py-2 text-white
            shadow-lg drop-shadow-sm hover:bg-opacity-90">Sign up
    </button>
</form>
