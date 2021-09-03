<template>
    <div>{{ this.sampleQuestionData }}</div>
</template>

<script>
export default {
    name: "AnswerTheQuestionPage",
    data() {
        return {
            sampleQuestionDatas: require('@/testdata/question.json')
        }
    },
    created() {
        let samples = [this.sampleQuestionDatas.q_4taku, this.sampleQuestionDatas.q_anaume, this.sampleQuestionDatas.q_collection]
        this.sampleQuestionData = samples[Math.floor(Math.random() * samples.length)];
    }
}
</script>