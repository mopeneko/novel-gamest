<template>
  <v-app>
    <v-navigation-drawer app v-model="isDrawerEnabled"></v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click.stop="isDrawerEnabled = !isDrawerEnabled"></v-app-bar-nav-icon>
      <v-toolbar-title>NovelGamest</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container fluid>
        <nuxt />
      </v-container>
    </v-content>

    <v-footer app></v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data() {
    return {
      isDrawerEnabled: false
    }
  }
})
</script>
