<template>
<div class="_player-selector">
  <PlaybackControls class="playback-controls" />

  <draft-seat-component
      v-for="seat in draft.seats"
      :key="seat.position"
      :seat="seat"
      />
</div>
</template>

<script lang="ts">
import Vue from 'vue';
import DraftSeatComponent from './DraftSeat.vue'
import PlaybackControls from './PlaybackControls.vue';
import { replayStore } from '../../state/ReplayStore';
import { draftStore } from '../../state/DraftStore';

import { DraftState } from '../../draft/DraftState';
import { navTo } from '../../router/url_manipulation';


export default Vue.extend({
  components: {
    DraftSeatComponent,
    PlaybackControls,
  },

  computed: {
    draft(): DraftState {
      return replayStore.draft;
    },
  },

  methods: {
    onNextClick() {
      replayStore.goNext();
      navTo(draftStore, replayStore, this.$route, this.$router, {});
    },

    onPrevClick() {
      replayStore.goBack();
      navTo(draftStore, replayStore, this.$route, this.$router, {});
    },

    onStartClick() {
      navTo(draftStore, replayStore, this.$route, this.$router, {
        eventIndex: 0,
      });
    },

    onEndClick() {
      navTo(draftStore, replayStore, this.$route, this.$router, {
        eventIndex: replayStore.events.length,
      });
    },
  },
});
</script>

<style scoped>
._player-selector {
  user-select: none;
  cursor: default;

  display: flex;
  flex-direction: column;
  overflow-y: auto;

  border-right: 1px solid #EAEAEA;
}

.playback-controls {
  padding: 10px 5px;
}
</style>
