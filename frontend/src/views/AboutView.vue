<template>
  <DefaultLayout>
    <div class="about px-6 md:px-12 lg:px-24">
      <div class="about-section">
        <h1 class="text-3xl md:text-5xl font-bold mb-6">About This Project</h1>
        <p class="mb-6 text-lg leading-relaxed">
          This application is a full-stack system built with <strong>Golang</strong> for the backend and <strong>Vue 3
            with TypeScript</strong> for the frontend. It retrieves stock recommendations from a third-party API, stores
          them in a <strong>CockroachDB</strong> database, and provides a clean, modern user interface where users can
          explore, filter, and analyze stocks.
        </p>
        <p class="mb-6 text-lg leading-relaxed">
          The backend is structured following <strong>Clean Architecture</strong> principles, using repositories, use
          cases, and service layers to ensure separation of concerns, maintainability, and testability. The frontend is
          organized with an <strong>Atomic Design</strong> approach, promoting scalability and component reusability.
        </p>
      </div>

      <div>
        <div class="about-section">
          <h2 class="text-2xl font-semibold mt-10 mb-4">About the Developer</h2>
          <p class="text-lg leading-relaxed">
            This project was designed and developed by a passionate full-stack developer with a strong focus on backend
            development, DevOps, and solution architecture. With solid experience in <strong>TypeScript</strong>,
            <strong>Python</strong>, and <strong>cloud-native technologies</strong>, the developer is committed to
            building
            scalable, maintainable, and production-ready systems.
          </p>
          <p class="mt-4 text-lg leading-relaxed">
            Key interests include <strong>SOLID principles</strong>, <strong>TDD</strong>, <strong>design
              patterns</strong>,
            and technologies like <strong>IaC (Infrastructure as Code)</strong>, continuous integration/delivery
            pipelines,
            and system observability. The vision is to continuously evolve toward becoming an expert in backend
            engineering
            and modern software architecture.
          </p>
        </div>
      </div>
    </div>
  </DefaultLayout>

</template>

<style>
.about {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.about-section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
</style>

<script setup lang="ts">
import DefaultLayout from '@/components/layouts/DefaultLayout.vue'
</script>