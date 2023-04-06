<template>
    <div>
        <header>
            <nav>
                <NuxtLink to="/">Hi</NuxtLink>
                <ul>
                    <li><NuxtLink to="/">Home</NuxtLink></li>
                    <li><NuxtLink to="/about">About</NuxtLink></li>
                    <li><NuxtLink to="/products">Products</NuxtLink></li>
                </ul>
                <a href="/about">Normal link</a>
            </nav>
        </header>

        <div>
            <slot />
        </div>
    </div>
</template>