# Install

```
go mod init test
kubebuilder init --domain test.org --license apache2 --owner "The Kubernetes demo"
kubebuilder create api --group test  --version v1beta1 --kind Test 
kubebuilder create webhook --group test --version v1beta1 --kind Test --defaulting --programmatic-validation
```

# 参考资料

> https://blog.51cto.com/u_13630803/2154192
> http://www.ansible.com.cn/docs/playbooks_intro.html
> https://www.cnblogs.com/wxzhe/p/10386649.html

# 架构

> playbook -》 pipeline -》 task -》actions and template and handlers(webhook or notify)

* playbook playbookrun
* pipeline pipelinerun
* task taskrun
* action
* template
* handlers is part of action
  * webhook
  * notify

# Init

```shell
go mod init hotwave
kubebuilder init --domain hotwave.io --license mit --owner "lflxp for China"
kubebuilder create api --group devops --version v1alpha1 --kind Playbook
kubebuilder create api --group devops --version v1alpha1 --kind Pipeline
kubebuilder create api --group devops --version v1alpha1 --kind PipelineRun
kubebuilder create api --group devops --version v1alpha1 --kind PlaybookRun
kubebuilder create api --group devops --version v1alpha1 --kind Task
kubebuilder create api --group devops --version v1alpha1 --kind TaskRun
kubebuilder create api --group devops --version v1alpha1 --kind Action
kubebuilder create api --group devops --version v1alpha1 --kind Template
```

# version

* v1alpha1
* v1beta1
* v1
