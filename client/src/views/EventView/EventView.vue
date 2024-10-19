<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { watch } from 'vue';

import Bets from './components/Bets.vue';
import Comments from './components/Comments.vue';
import Conditions from './components/Conditions.vue';
import Graph from './components/Graph.vue';
import Information from './components/Information.vue';

import { type Event, useEvents } from '@/services/events';

const events = useEvents();

const { eventID } = defineProps<{
	eventID: Event['id'];
}>();

events.select(eventID);
watch(
	() => eventID,
	value => {
		events.select(value);
	},
);
const { current: event } = storeToRefs(events);
</script>

<template>
	<div
		class="event"
		v-if="event"
	>
		<Graph :event="event" />
		<Bets :event="event" />
		<Information />
		<Conditions />
		<Comments />
	</div>
	<div v-else>Event was not found!</div>
</template>

<style>
.event {
	display: flex;
	flex-direction: column;
	gap: 6px;
}
</style>
