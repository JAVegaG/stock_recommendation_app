<template>
    <div class="min-h-screen bg-gray-50">
        <nav class="bg-white p-4 shadow">
            <slot name="nav" />
        </nav>
        <main class="p-6">
            <slot />
        </main>
    </div>
</template>

<script setup lang="ts" />
