# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.define "oc-server"
  config.vm.box = "ubuntu/trusty64"

  config.vm.network "private_network", ip: "192.168.128.100"
  config.vm.hostname = "orderchef"

  config.vm.synced_folder ".",
    "/orderchef/src/lab.castawaylabs.com/orderchef",
    type: "rsync", rsync__exclude: [".vagrant", ".git"]

  config.vm.provider "virtualbox" do |vb|
    vb.cpus = 2
    vb.memory = 2048
    vb.name = "orderchef-server"
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible/orderchef.yml"
  end
end
