<script lang="ts">
	import { goto } from '$app/navigation';
	import DarkMode from '$components/dark-mode.svelte';
	import Head from '$components/head.svelte';
	import Icon from '$components/icon.svelte';
	import Rooms from '$components/room/rooms.svelte';
	import { DEFAULT_ROOM, PROFILE_PATH, ROOM_PATH } from '$lib/constant';
	import { currentRoom, refreshRooms, roomStore } from '$lib/store';
	import {
		modalStore,
		type ModalComponent,
		type ModalSettings,
		AppShell,
		AppBar
	} from '@skeletonlabs/skeleton';

	let loaded = false;
	let currentRoomTitle = '';
	// get current room information
	$: if (!$roomStore.loading && $roomStore.data.length != 0 && !loaded) {
		loaded = true;
		currentRoomTitle = $roomStore.data.find((e) => e.slug == $currentRoom)?.name || DEFAULT_ROOM;
	}
</script>

<Head />

<AppShell>
	<svelte:fragment slot="header">
		<AppBar>
			<svelte:fragment slot="lead">
				<a href={ROOM_PATH + '/' + $currentRoom} class="btn-icon btn-icon-sm variant-soft">
					<Icon name="arrow-left" width="18px" height="18px" />
				</a>
			</svelte:fragment>
			<a href={ROOM_PATH + '/' + $currentRoom}>Go to {currentRoomTitle}</a>
			<svelte:fragment slot="trail">
				<DarkMode />
				<a class="btn-icon btn-icon-sm variant-soft" href={PROFILE_PATH} title="Profile">
					<Icon name="user" width="18px" height="18px" />
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- (sidebarLeft) -->
	<!-- (sidebarRight) -->
	<!-- (pageHeader) -->
	<!-- Router Slot -->
	<div class="card max-w-[640px] mx-auto my-4">
		<div class="w-full inline-flex justify-between items-center p-4 pb-0">
			<h1 class="h3 font-bold">Chat Rooms</h1>
			<Icon name="cat" class="text-purple-600" width="24px" height="24px" />
		</div>
		<Rooms />
	</div>
	<!-- ---- / ---- -->
	<!-- (pageFooter) -->
	<!-- (footer) -->
</AppShell>
