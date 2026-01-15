<script setup lang="ts">

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Card from 'primevue/card';

defineProps<{ }>()

let scores = [
  { rank: 1, username: 'PlayerOne', score: 1500 },
  { rank: 2, username: 'GamerGal', score: 1450 },
  { rank: 3, username: 'NoobMaster', score: 1400 },
  { rank: 4, username: 'ProGamer', score: 1350 },
  { rank: 5, username: 'Speedy', score: 1300 },
  { rank: 1, username: 'PlayerOne', score: 1500 },
  { rank: 2, username: 'GamerGal', score: 1450 },
  { rank: 3, username: 'NoobMaster', score: 1400 },
  { rank: 4, username: 'ProGamer', score: 1350 },
  { rank: 5, username: 'Speedy', score: 1300 },
  { rank: 1, username: 'PlayerOne', score: 1500 },
];


</script>

<template>
  <Card>
    <template #content>
      <DataTable class="bg-heroic m-10 border-2 border-white"
        :value="scores"
        showGridlines
      >
        <Column field="rank" header="Rank"></Column>
        <Column field="username" header="Username"></Column>
        <Column field="score" header="Score"></Column>
      </DataTable>
    </template>
  </Card>
</template>