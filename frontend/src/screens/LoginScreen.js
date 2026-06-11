import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, StyleSheet } from 'react-native';
import { colors, radius, spacing } from '../constants/theme';

export default function LoginScreen({ navigation }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = () => {
    // TODO (Sprint 2): llamar a /api/v1/auth/login y guardar el token
    console.log('Login:', { email, password });
  };

  return (
    <View style={styles.container}>
      <View style={styles.logoBlock}>
        <View style={styles.logoIcon} />
        <Text style={styles.logoText}>Gymify</Text>
        <Text style={styles.tagline}>Entrena. Comparte. Crece.</Text>
      </View>

      <Text style={styles.label}>Correo electrónico</Text>
      <TextInput
        style={styles.input}
        placeholder="tu@correo.com"
        placeholderTextColor={colors.textMuted}
        value={email}
        onChangeText={setEmail}
        autoCapitalize="none"
        keyboardType="email-address"
      />

      <Text style={styles.label}>Contraseña</Text>
      <TextInput
        style={styles.input}
        placeholder="••••••••"
        placeholderTextColor={colors.textMuted}
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />

      <TouchableOpacity onPress={() => {/* TODO: recuperación de contraseña */}}>
        <Text style={styles.forgot}>¿Olvidaste tu contraseña?</Text>
      </TouchableOpacity>

      <TouchableOpacity style={styles.button} onPress={handleLogin}>
        <Text style={styles.buttonText}>Entrar</Text>
      </TouchableOpacity>

      <TouchableOpacity onPress={() => navigation?.navigate?.('Register')}>
        <Text style={styles.registerLink}>
          ¿No tienes cuenta? <Text style={styles.registerLinkAccent}>Regístrate</Text>
        </Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
    paddingHorizontal: spacing.xl,
    paddingTop: 60,
  },
  logoBlock: {
    alignItems: 'center',
    marginBottom: spacing.xl * 1.5,
  },
  logoIcon: {
    width: 56,
    height: 56,
    borderRadius: radius.lg,
    backgroundColor: colors.accentSoft,
    borderWidth: 1.5,
    borderColor: colors.accent,
    marginBottom: spacing.md,
  },
  logoText: {
    color: colors.text,
    fontSize: 28,
    fontWeight: '900',
    letterSpacing: -1,
  },
  tagline: {
    color: colors.textMuted,
    fontSize: 11,
    marginTop: spacing.xs,
    letterSpacing: 1.5,
    textTransform: 'uppercase',
  },
  label: {
    color: colors.text,
    fontSize: 12,
    fontWeight: '700',
    marginBottom: spacing.xs,
  },
  input: {
    backgroundColor: colors.card2,
    borderWidth: 0.5,
    borderColor: colors.border,
    borderRadius: radius.md,
    paddingHorizontal: spacing.md,
    paddingVertical: spacing.md,
    color: colors.text,
    marginBottom: spacing.md,
  },
  forgot: {
    color: colors.accent,
    fontSize: 11,
    textAlign: 'right',
    marginBottom: spacing.lg,
  },
  button: {
    backgroundColor: colors.accent,
    borderRadius: radius.md,
    paddingVertical: spacing.md,
    alignItems: 'center',
    marginBottom: spacing.lg,
  },
  buttonText: {
    color: '#fff',
    fontWeight: '800',
    fontSize: 14,
  },
  registerLink: {
    color: colors.textMuted,
    fontSize: 12,
    textAlign: 'center',
  },
  registerLinkAccent: {
    color: colors.accent,
    fontWeight: '700',
  },
});
