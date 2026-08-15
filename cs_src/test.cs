using System;
using System.Collections.Generic;

namespace GoSynapseDemo
{
    public class SecurityAuditor
    {
        private List<string> _auditLogs = new List<string>();

        public void AuditAccess(string principal, string resource)
        {
            Console.WriteLine($"Auditing access: {principal} -> {resource}");
            LogEvent($"{principal} requested access to {resource}");
        }

        private void LogEvent(string detail)
        {
            _auditLogs.Add($"[{DateTime.UtcNow:o}] {detail}");
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var auditor = new SecurityAuditor();
            auditor.AuditAccess("admin_user", "customer_records");
        }
    }
}
