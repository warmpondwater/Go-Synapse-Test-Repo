using System;
using System.Collections.Generic;

namespace GoSynapseDemo
{
    public interface ISecurityAuditor
    {
        void AuditAccess(string principal, string resource);
        bool IsHealthy();
    }

    // (DEAD CODE)
    public class UnusedCSharpService
    {
        public void AbandonedOperation()
        {
            Console.WriteLine("Dead code in C#");
        }
    }

    public class SecurityAuditor : ISecurityAuditor
    {
        private List<string> _auditLogs = new List<string>();
        private int _totalEvents = 0;

        // (SOURCE)
        public string GetUntrustedPrincipal()
        {
            return "user_injection_candidate_admin";
        }

        // (SANITIZER)
        public string SanitizePrincipal(string raw)
        {
            return raw.Replace("'", "").Replace("\"", "");
        }

        // (SINK)
        public void DispatchToSecurityLog(string sanitizedPrincipal, string resource)
        {
            Console.WriteLine($"[C# SINK LOG]: {sanitizedPrincipal} accessed {resource}");
        }

        public void AuditAccess(string principal, string resource)
        {
            string clean = SanitizePrincipal(principal);
            DispatchToSecurityLog(clean, resource);
            LogEvent($"{clean} requested access to {resource}");
        }

        public bool IsHealthy()
        {
            return _totalEvents >= 0;
        }

        private void LogEvent(string detail)
        {
            _totalEvents++;
            _auditLogs.Add($"[{DateTime.UtcNow:o}] {detail}");
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            ISecurityAuditor auditor = new SecurityAuditor();
            var concrete = new SecurityAuditor();
            string rawUser = concrete.GetUntrustedPrincipal();
            auditor.AuditAccess(rawUser, "customer_records");
        }
    }
}

