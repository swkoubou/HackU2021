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
.syomons {
  align-items: center;
  justify-content: center;
}
.syomon {
  background: rgb(236, 236, 236);
  width: 80%;
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 5px;
}
</style>
