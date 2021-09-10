<template>
  <div>
    <h2>スコア</h2>
    <!-- <h4>{{answersData}}</h4> -->
    <div class="daimons">
      <div
        class="daimon"
        v-for="(daimon, daimonIndex) in answersData"
        :key="`daimon-${daimonIndex}`"
      >
        <div class="question-score-card">
          <h3 class="daimon-no">大問{{ daimonIndex + 1 }}</h3>
          <div class="syomons">
            <div
              class="syomon"
              v-for="(syomon, syomonIndex) in daimon['questionAnswers']"
              :key="`syomon-${syomonIndex}`"
            >
              <ScoreSyomon
                :SyomonNo="syomonIndex + 1"
                :UserAnswer="daimon['userAnswers'][syomonIndex]"
                :QuestionAnswer="syomon"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ScoreSyomon from '@/components/ScoreSyomon.vue'

export default {
  name: 'ScorePage',
  props: {
    answersData: {
      type: Array,
      required: true,
    },
  },
  components: {
    ScoreSyomon,
  },
}
</script>

<style scoped>
.syomon {
  background: rgb(236, 236, 236);
  max-width: 800px;
  margin: 5px 0px;
  margin-left: auto;
  margin-right: auto;
}

/* 横幅が少ない時の最低マージン */
@media (max-width: 840px) {
  .syomon {
    margin-left: 20px;
    margin-right: 20px;
  }
}
</style>
