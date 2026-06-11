import React from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { colors, radius, spacing } from '../constants/theme';

// TODO (Sprint 3): reemplazar datos de ejemplo por datos reales del backend
// (GET /api/v1/streaks/user/{id})
const MOCK_USER = {
  name: 'Bianca',
  streak: 14,
  level: 7,
  xp: 2340,
  xpToNextLevel: 660,
  weekProgress: ['done', 'done', 'done', 'today', 'pending', 'pending', 'pending'],
};

const WEEK_LABELS = ['L', 'M', 'X', 'J', 'V', 'S', 'D'];

export default function HomeScreen() {
  const xpProgress = 78; // porcentaje hacia el siguiente nivel

  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <View style={styles.header}>
        <View>
          <Text style={styles.logo}>Gymify</Text>
          <Text style={styles.greeting}>Buenos días, {MOCK_USER.name}</Text>
        </View>
      </View>

      {/* Racha + nivel */}
      <View style={styles.card}>
        <View style={styles.streakRow}>
          <View>
            <Text style={styles.cardLabel}>Racha activa</Text>
            <Text style={styles.streakNumber}>{MOCK_USER.streak}</Text>
            <Text style={styles.streakSub}>días consecutivos</Text>
          </View>
          <View style={styles.levelBlock}>
            <View style={styles.levelPill}>
              <Text style={styles.levelPillText}>NIVEL {MOCK_USER.level}</Text>
            </View>
            <Text style={styles.xpText}>{MOCK_USER.xp.toLocaleString()} XP</Text>
            <Text style={styles.xpSub}>{MOCK_USER.xpToNextLevel} para nivel {MOCK_USER.level + 1}</Text>
          </View>
        </View>
        <View style={styles.progressBarBg}>
          <View style={[styles.progressBarFill, { width: `${xpProgress}%` }]} />
        </View>
      </View>

      {/* Semana */}
      <View style={styles.card}>
        <Text style={styles.cardLabelSmall}>Esta semana</Text>
        <View style={styles.weekRow}>
          {WEEK_LABELS.map((label, i) => {
            const status = MOCK_USER.weekProgress[i];
            return (
              <View key={label} style={styles.dayColumn}>
                <View
                  style={[
                    styles.dayCircle,
                    status === 'done' && styles.dayCircleDone,
                    status === 'today' && styles.dayCircleToday,
                  ]}
                >
                  <Text
                    style={[
                      styles.dayCircleText,
                      status === 'done' && styles.dayCircleTextDone,
                      status === 'today' && styles.dayCircleTextToday,
                    ]}
                  >
                    {label}
                  </Text>
                </View>
                {status === 'today' && <Text style={styles.todayLabel}>hoy</Text>}
              </View>
            );
          })}
        </View>
      </View>

      {/* Reto activo */}
      <Text style={styles.sectionTitle}>Reto activo</Text>
      <View style={styles.card}>
        <View style={styles.challengeHeader}>
          <Text style={styles.challengeTitle}>30 días sin falta</Text>
          <View style={styles.pill}>
            <Text style={styles.pillText}>14/30</Text>
          </View>
        </View>
        <Text style={styles.challengeSub}>23 participantes · Vas en 2do lugar</Text>
        <View style={styles.progressBarBg}>
          <View style={[styles.progressBarFill, { width: '47%', backgroundColor: colors.xp }]} />
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
  },
  content: {
    padding: spacing.md,
    paddingTop: 50,
  },
  header: {
    marginBottom: spacing.md,
  },
  logo: {
    color: colors.accent,
    fontSize: 17,
    fontWeight: '800',
    letterSpacing: -1,
  },
  greeting: {
    color: colors.textMuted,
    fontSize: 11,
    marginTop: 2,
  },
  card: {
    backgroundColor: colors.card,
    borderWidth: 0.5,
    borderColor: colors.border,
    borderRadius: radius.lg,
    padding: spacing.md,
    marginBottom: spacing.sm,
  },
  cardLabel: {
    color: colors.textMuted,
    fontSize: 10,
    textTransform: 'uppercase',
    letterSpacing: 1,
    marginBottom: 2,
  },
  cardLabelSmall: {
    color: colors.textMuted,
    fontSize: 10,
    textTransform: 'uppercase',
    letterSpacing: 1,
    marginBottom: spacing.sm,
  },
  streakRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    marginBottom: spacing.sm,
  },
  streakNumber: {
    color: colors.text,
    fontSize: 32,
    fontWeight: '900',
    letterSpacing: -1,
    lineHeight: 36,
  },
  streakSub: {
    color: colors.textMuted,
    fontSize: 10,
  },
  levelBlock: {
    alignItems: 'flex-end',
  },
  levelPill: {
    backgroundColor: colors.accentSoft,
    borderWidth: 0.5,
    borderColor: colors.accent,
    borderRadius: radius.pill,
    paddingHorizontal: spacing.sm,
    paddingVertical: 2,
    marginBottom: 6,
  },
  levelPillText: {
    color: colors.accent,
    fontSize: 9,
    fontWeight: '700',
  },
  xpText: {
    color: colors.xp,
    fontSize: 13,
    fontWeight: '800',
  },
  xpSub: {
    color: colors.textMuted,
    fontSize: 9,
  },
  progressBarBg: {
    backgroundColor: '#222',
    borderRadius: 4,
    height: 5,
  },
  progressBarFill: {
    backgroundColor: colors.accent,
    borderRadius: 4,
    height: 5,
  },
  weekRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  dayColumn: {
    alignItems: 'center',
  },
  dayCircle: {
    width: 34,
    height: 34,
    borderRadius: 17,
    borderWidth: 1.5,
    borderColor: colors.border,
    alignItems: 'center',
    justifyContent: 'center',
  },
  dayCircleDone: {
    backgroundColor: colors.accent,
    borderColor: colors.accent,
  },
  dayCircleToday: {
    borderColor: colors.accent,
  },
  dayCircleText: {
    color: colors.textMuted,
    fontSize: 11,
    fontWeight: '700',
  },
  dayCircleTextDone: {
    color: '#fff',
  },
  dayCircleTextToday: {
    color: colors.accent,
  },
  todayLabel: {
    color: colors.accent,
    fontSize: 8,
    marginTop: 2,
  },
  sectionTitle: {
    color: colors.text,
    fontSize: 12,
    fontWeight: '700',
    marginBottom: spacing.sm,
    marginTop: spacing.xs,
  },
  challengeHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 4,
  },
  challengeTitle: {
    color: colors.text,
    fontSize: 13,
    fontWeight: '700',
  },
  challengeSub: {
    color: colors.textMuted,
    fontSize: 10,
    marginBottom: spacing.sm,
  },
  pill: {
    backgroundColor: colors.accentSoft,
    borderWidth: 0.5,
    borderColor: colors.accent,
    borderRadius: radius.pill,
    paddingHorizontal: spacing.sm,
    paddingVertical: 2,
  },
  pillText: {
    color: colors.accent,
    fontSize: 10,
    fontWeight: '700',
  },
});
