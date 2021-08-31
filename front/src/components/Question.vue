<template>
    <div class="question-card" id="question-parent">
        <div v-if="questionData.questionType === '4taku'">
            <Question4taku :question="questionData"/>
            </div>
        <div v-else-if="questionData.questionType === 'anaume'">
            <QuestionAnaume :question="questionData"/>
        </div>
        <div v-else>
            <h3>
                unknown : {{ questionData.questionType }}
            </h3>
        </div>
    </div>
</template>

<script>
import Question4taku from "@/components/Question4taku.vue";
import QuestionAnaume from "@/components/QuestionAnaume.vue";

export default {
    name: "QuestionParent",
    props: ["questionData"],
    components: {
        Question4taku,
        QuestionAnaume
    },
    setup() {
    },
}
</script>

<style scoped>
.question-card {
    background: #e2e1e0;
    border-radius: 10px;
    display: inline-block;
    height: 500px;
    margin: 1rem;
    position: relative;
    width: 500px;
}
</style>
